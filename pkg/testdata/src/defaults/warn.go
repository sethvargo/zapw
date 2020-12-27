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

func Warn_ok() {
	logger.Warn("goodbye", "cruel")
}

func Warnf_ok() {
	logger.Warnf("goodbye", "cruel")
}

func Warnw_ok() {
	logger.Warnw("goodbye")
}

func Warnw_ok_args() {
	logger.Warnw("goodbye", "cruel", "world")
}

func Warnw_fail_numargs() {
	logger.Warnw("goodbye", "cruel") // want `zap.SugaredLogger must have an even number .+`
}

func Warnw_fail_typeargs() {
	logger.Warnw("goodbye", 64, "universe") // want `zap.SugaredLogger requires keys to be strings \(got int\)`
}
