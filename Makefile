# https://gist.github.com/alexedwards/3b40775846535d0014ab1ff477e4a568

# =====================
#       VARIABLES
# =====================
BUILD_PATH := ./bin

BINARY_NAME := timer
PACKAGE_PATH := ./cmd/cli

# =====================
#      DEVELOPMENT
# =====================
## build: build cli app
.PHONY: build
build:
	go build \
    -ldflags " \
      -X timer/internal/version.AppVersion=$(shell git tag | tail -n 1 | sed 's/^v//') \
      -X \"timer/internal/version.BuildTime=$(shell date '+%d.%m.%Y %H:%M:%S')\" \
      -X timer/internal/version.CommitHash=$(shell git rev-parse --short HEAD)" \
    -o=${BUILD_PATH}/${BINARY_NAME} \
    ${PACKAGE_PATH}/main.go

## run: run command line application
.PHONY: run
run: build
	${BUILD_PATH}/${BINARY_NAME} ${ARGS}

## test: run all tests
.PHONY: test
test:
	go test -cover ./...

## test/cover: run all tests and display coverage
.PHONY: test/cover
test/cover:
	go test -v -race -coverprofile=./coverage.out ./...
	go tool cover -html=./coverage.out


# =====================
#    QUALITY CONTROL
# =====================
## tidy: format code and tidy modfile
.PHONY: tidy
tidy:
	go fmt ./...
	go mod tidy -v

## audit: run quality control checks
.PHONY: audit
audit:
	go mod verify


# =====================
#        HELPERS
# =====================
## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

## init/hooks: specify the path for git hooks
.PHONY: init/hooks
init/hooks:
	git config core.hooksPath .github/hooks

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]
