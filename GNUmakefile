.PHONY: fun-tests unit-tests test

export

project_name := travis-go-test
GORELEASER_VERSION := 0.179.0
LINTER_VERSION := 1.43.0
PATH := $(shell pwd)/bin:$(PATH)
SHELL := bash

.ONESHELL:
.SHELLFLAGS := -ec

bash-required-version:
	$(if $(filter oneshell,${.FEATURES}) \
	  , \
	  ,$(error oneshell not supported - update your make))

check-bash: bash-required-version
	@test $$BASH_VERSINFO = 5 || test $$BASH_VERSINFO = 4

bin: check-bash
	@mkdir -p bin/

bin/golangci-lint: bin
	@test -f $@ ||
	(
		curl -fsL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v${LINTER_VERSION}
		bin/golangci-lint version
	)

bin/goreleaser: bin
	@test -f $@ ||
	(
		curl -fsL https://github.com/goreleaser/goreleaser/releases/download/v${GORELEASER_VERSION}/goreleaser_`uname`_x86_64.tar.gz -o goreleaser.tgz
		tar xvf goreleaser.tgz -C bin goreleaser
		rm goreleaser.tgz
		bin/goreleaser --version
	)

tools: bin/goreleaser bin/golangci-lint

pre-check: tools
	[[ `(gofmt -l .)`x == x ]] || (echo "go fmt failed" && gofmt -l . && exit 1)
	go vet ./...
	golangci-lint run ./...

release-test: build
	@echo $@ passed
	#bin/vergo check increment-hint

release: build
	@echo $@ passed && exit 0
	bin/vergo check release || exit 0
	bin/vergo bump auto
	BUILT_BY="`goreleaser --version | head -n1`, `go version`" \
	GORELEASER_CURRENT_TAG=`bin/vergo get latest-release -p` \
	GORELEASER_PREVIOUS_TAG=`bin/vergo get previous-release -p` \
	goreleaser release --skip-validate --rm-dist
	bin/vergo push

unit-tests: pre-check
	go clean -testcache
	go test ./...

fun-tests: build
	@echo $@ passed

test: build unit-tests fun-tests
test-compile:
	go test --exec=true ./...

build: pre-check
	BUILT_BY="`goreleaser --version | head -n1`, `go version`" \
	GORELEASER_CURRENT_TAG=0+`git rev-parse --short HEAD` \
	goreleaser build --snapshot --rm-dist
	@dist/${project_name}_`uname | tr A-Z a-z`_amd64/${project_name} version
	@cp dist/${project_name}_`uname | tr A-Z a-z`_amd64/${project_name} bin/${project_name}
	@cp dist/${project_name}_`uname | tr A-Z a-z`_amd64/${project_name} /usr/local/bin/${project_name} 2>/dev/null || true

dependency-updates:
	@go list -u -f '{{if (and (not (or .Main .Indirect)) .Update)}}{{.Path}}: {{.Version}} -> {{.Update.Version}}{{end}}' -m all

go-licenses:
	@go install github.com/google/go-licenses@latest

print-licenses: go-licenses
	@go-licenses csv .

check-licenses: go-licenses
	!(go-licenses csv . | grep -E 'GNU|AGPL|GPL|MPL|CPL|CDDL|EPL|CCBYNC|Facebook|WTFPL')