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

package defaults

type customString string
type stringAlias = string

const constCustomString = customString("cruel")
const constStringAlias = stringAlias("cruel")

var varCustomString = customString("cruel")
var varStringAlias = stringAlias("cruel")

func customString_ok() {
	logger.Errorw("goodbye", customString("cruel"), "world")
}

func stringAlias_ok() {
	logger.Errorw("goodbye", stringAlias("cruel"), "world")
}

func constCustomString_ok() {
	logger.Errorw("goodbye", constCustomString, "world")
}

func constStringAlias_ok() {
	logger.Errorw("goodbye", constStringAlias, "world")
}

func varCustomString_ok() {
	logger.Errorw("goodbye", varCustomString, "world")
}

func varStringAlias_ok() {
	logger.Errorw("goodbye", varStringAlias, "world")
}

type customInt int
type intAlias = int

const constCustomInt = customInt(42)
const constIntAlias = intAlias(42)

var varCustomInt = customInt(42)
var varIntAlias = intAlias(42)

func customInt_ok() {
	logger.Errorw("goodbye", customInt(42), "world") // want `zap.SugaredLogger requires keys to be strings .+`
}

func intAlias_ok() {
	logger.Errorw("goodbye", intAlias(42), "world") // want `zap.SugaredLogger requires keys to be strings .+`
}

func constCustomInt_ok() {
	logger.Errorw("goodbye", constCustomInt, "world") // want `zap.SugaredLogger requires keys to be strings .+`
}

func constIntAlias_ok() {
	logger.Errorw("goodbye", constIntAlias, "world") // want `zap.SugaredLogger requires keys to be strings .+`
}

func varCustomInt_ok() {
	logger.Errorw("goodbye", varCustomInt, "world") // want `zap.SugaredLogger requires keys to be strings .+`
}

func varIntAlias_ok() {
	logger.Errorw("goodbye", varIntAlias, "world") // want `zap.SugaredLogger requires keys to be strings .+`
}
