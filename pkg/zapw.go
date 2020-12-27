// Copyright 2020 Seth Vargo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package zapw defines an analyzer that finds mistakes with variadic "with"
// arguments in zap's SugaredLogger.
package zapw

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"go/types"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

// Analyzer is the root analyzer.
var Analyzer = &analysis.Analyzer{
	Name: "zapw",
	Doc:  "check for invalid variadic arguments to zap.SugaredLogger",
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

var (
	// zapWithFuncName is the name of the zap function "With".
	zapWithFuncName = "With"

	// zapWFuncNames is the list of function names that accept variadic key-value
	// pairs.
	zapWFuncNames = map[string]struct{}{
		"DPanicw":       {},
		"Debugw":        {},
		"Errorw":        {},
		"Fatalw":        {},
		"Infow":         {},
		"Panicw":        {},
		"Warnw":         {},
		zapWithFuncName: {},
	}

	// zapPackageImportPath is the zap package import path.
	zapPackageImportPath = "go.uber.org/zap"
)

const (
	// zapSugaredLoggerIdent is the string ident of zap's sugared logger.
	zapSugaredLoggerIdent = "go.uber.org/zap.SugaredLogger"
)

// run is the main entrypoint for the analyzer.
func run(pass *analysis.Pass) (interface{}, error) {
	// Do not process zap itself.
	if pass.Pkg.Path() == zapPackageImportPath {
		return nil, nil
	}

	// Build the inspector.
	inspector, ok := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	if !ok {
		return nil, fmt.Errorf("%T is not *inspector.Inspector", inspector)
	}

	filter := []ast.Node{&ast.CallExpr{}}
	inspector.Preorder(filter, func(n ast.Node) {
		if n == nil {
			return
		}

		// This should never happen, but better safe than panic.
		callExpr, ok := n.(*ast.CallExpr)
		if !ok {
			return
		}

		// Pull out the selector and selection so we can process it.
		if callExpr.Fun == nil {
			return
		}
		selectorExpr, ok := callExpr.Fun.(*ast.SelectorExpr)
		if !ok {
			return
		}

		selection, ok := pass.TypesInfo.Selections[selectorExpr]
		if !ok {
			return
		}

		selectionObj := selection.Obj()
		if selectionObj == nil {
			return
		}

		// Determine if the selection could be a sugared logger. Right now, this
		// only checks if the receiver is a logger.
		if !isZapSugaredLogger(pass, selectionObj) {
			return
		}

		// Get the arguments.
		args := callExpr.Args
		if selectionObj.Name() != zapWithFuncName {
			// Regular "w" functions expect the first parameter to be a string, then
			// the remainder are variadic. Since the function definition requires the
			// first parameter to be a string, this is safe to without a bounds check.
			args = args[1:]
		}

		// Ensure there are an even number of arguments.
		if len(args)%2 != 0 {
			pass.ReportRangef(n, `zap.SugaredLogger must have an even number of "With" elements`)

			// If there's an odd number of arguments, checking individual arguments is
			// probably going to return a bunch of unnecessary linter errors, so
			// return now. After the user fixes the number of elements, re-running the
			// linter will check the argument types.
			return
		}

		// Verify that elements at even indicies are string types.
		for i := 0; i < len(args); i += 2 {
			arg := args[i]
			typ := pass.TypesInfo.TypeOf(arg)

			basic, ok := toBasicType(typ)
			if !ok || basic.Kind() != types.String {
				pass.ReportRangef(arg, fmt.Sprintf(`zap.SugaredLogger requires keys to be strings (got %s)`, typ))
			}
		}
	})

	return nil, nil
}

// isZapSugaredLogger checks if the receiver is a go.uber.org/zap.SugaredLogger.
func isZapSugaredLogger(pass *analysis.Pass, obj types.Object) bool {
	// There's no point in doing all the reflection if the method name isn't one
	// we're searching for.
	if _, ok := zapWFuncNames[obj.Name()]; !ok {
		return false
	}

	if !strings.Contains(obj.String(), zapSugaredLoggerIdent) {
		return false
	}

	return true
}

// toBasicType attempts to get the underlying basic type. This is to handle edge
// cases where the developer has aliased or created a custom type, but we really
// only care about the basic/primitive type.
func toBasicType(typ types.Type) (*types.Basic, bool) {
	for i := 0; i < 10; i++ {
		if basic, ok := typ.(*types.Basic); ok {
			return basic, true
		}

		typ = typ.Underlying()
	}

	return nil, false
}

// derefTypesPointer dereferences the given type if it's a pointer. It bounds at
// 10 attempts before bailing to prevent an infinite loop.
func derefTypesPointer(typ types.Type) types.Type {
	for i := 0; i < 10; i++ {
		ptr, ok := typ.(*types.Pointer)
		if !ok {
			return typ
		}
		typ = ptr.Elem()
	}

	panic(fmt.Sprintf("pointer did not deref after 10 tries: %T", typ))
}

// printNode prints the ast.Node. It's used for debugging.
func printNode(n ast.Node) string {
	var b bytes.Buffer
	fset := token.NewFileSet()
	if err := printer.Fprint(&b, fset, n); err != nil {
		panic(err)
	}
	return b.String()
}
