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

func Debug_ok() {
	logger.Debug("goodbye", "cruel")
}

func Debugf_ok() {
	logger.Debugf("goodbye", "cruel")
}

func Debugw_ok() {
	logger.Debugw("goodbye")
}

func Debugw_ok_args() {
	logger.Debugw("goodbye", "cruel", "world")
}

func Debugw_fail_numargs() {
	logger.Debugw("goodbye", "cruel") // want `zap.SugaredLogger must have an even number .+`
}

func Debugw_fail_typeargs() {
	logger.Debugw("goodbye", 64, "universe") // want `zap.SugaredLogger requires keys to be strings \(got int\)`
}
