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

func Panic_ok() {
	logger.Panic("goodbye", "cruel")
}

func Panicf_ok() {
	logger.Panicf("goodbye", "cruel")
}

func Panicw_ok() {
	logger.Panicw("goodbye")
}

func Panicw_ok_args() {
	logger.Panicw("goodbye", "cruel", "world")
}

func Panicw_fail_numargs() {
	logger.Panicw("goodbye", "cruel") // want `zap.SugaredLogger must have an even number .+`
}

func Panicw_fail_typeargs() {
	logger.Panicw("goodbye", 64, "universe") // want `zap.SugaredLogger requires keys to be strings \(got int\)`
}
