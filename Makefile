# Copyright 2020 Seth Vargo
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

plugin:
	@go build -buildmode=plugin -o ./bin/zapw.so ./plugin/example.go
.PHONY: plugin

test: zap-source
	@go test \
		-count=1 \
		-timeout=5m \
		-race \
		./...
.PHONY: test

zap-source:
	@cd pkg/testdata/src && \
		rm -rf go.uber.org && \
		git clone --quiet --depth=1 https://github.com/uber-go/zap go.uber.org/zap && \
		git clone --quiet --depth=1 https://github.com/uber-go/atomic go.uber.org/atomic && \
		git clone --quiet --depth=1 https://github.com/uber-go/multierr go.uber.org/multierr
.PHONY: zap-source
