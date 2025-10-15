# Copyright 2020-2025 SIX AFTER, INC (SIX AFTER)
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

SHELL := /bin/bash

.DEFAULT: ;: do nothing
GO_CMD=go
GO_BUILD=$(GO_CMD) build
GO_CLEAN=$(GO_CMD) clean
GO_TEST=$(GO_CMD) test
GO_GET=$(GO_CMD) get
GO_VET=$(GO_CMD) vet
GO_FMT=$(GO_CMD) fmt
GO_MOD=$(GO_CMD) mod
GO_LINT_CMD=golangci-lint run
GO_WORK=$(GO_CMD) work
GO_WORK_FILE := ./go.work
FUZZTIME ?= 20s

export BINARY_NAME=main.out

.PHONY: build
build: ## Build the binary file
	@scripts/go-build.sh

.PHONY: test
test: ## Execute unit tests
	@scripts/go-test.sh

.PHONY: run
run:
	@scripts/go-run.sh

.PHONY: clean
clean: ## Remove previous build
	@scripts/go-clean.sh

.PHONY: cover
cover: ## Generate global code coverage report
	@scripts/go-cover.sh

.PHONY: analyze
analyze: ## Generate static analysis report
	@scripts/go-analyze.sh

.PHONY: deps
deps: ## Get the dependencies and vendor
	@scripts/go-deps.sh

.PHONY: fmt
fmt: ## Format the files
	$(GO_FMT) ./..

.PHONY: vet
vet: ## Vet the files
	$(GO_VET) -v ./...

.PHONY: lint
lint: ## Lint the files
	$(GO_LINT_CMD) --config .golangci.yaml --verbose ./...

.PHONY: mod-download
mod-download:
	$(GO_MOD) download

.PHONY: vendor
vendor:
	$(GO_MOD) vendor

.PHONY: tidy
tidy:
	$(GO_MOD) tidy

.PHONY: update
update:
	$(GO_GET) -u

.PHONY: vuln
vuln: ## Check for vulnerabilities
	govulncheck ./...

.PHONY: release-verify
release-verify: ## Verify the release
	rm -fr dist
	goreleaser --config .goreleaser.yaml release --snapshot

.PHONY: help
help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' | sort

# %: - rule which match any task name;  @: - empty recipe = do nothing
%:
    @:
