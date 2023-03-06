help: ## show this message
	@echo "All commands can be run on local machine as well as inside dev container."
	@echo ""
	@sed -nE 's/^ *([^[:blank:]]+)[[:blank:]]*:[^#]*##[[:blank:]]*(.+)/\1\n\2/p' $(MAKEFILE_LIST) | tr '\n' '\0' | xargs -0 -n 2 printf '%-25s%s\n'
.PHONY: help

.DEFAULT_GOAL := help

test: ## run all tests
	@echo "+ $@"
	go test -race -count 1 -p 8 -parallel 8 -timeout 1m ./...
.PHONY: test

test-cover: ## run all tests with code coverage
	@echo "+ $@"
	go test -race -count 1 -p 8 -parallel 8 -timeout 1m -coverpkg ./... -coverprofile coverage.out ./...
.PHONY: test-cover

lint: build-docker-dev ## run linter
	@echo "+ $@"
	$(RUN_IN_DOCKER) golangci-lint run
.PHONY: lint

bash: build-docker-dev ## run bash inside container for development
 ifndef INSIDE_DEV_CONTAINER
	@echo "+ $@"
	$(RUN_IN_DOCKER) bash
 endif
.PHONY: bash

check-tidy: ## ensure go.mod is tidy
	@echo "+ $@"
	cp go.mod go.check.mod
	cp go.sum go.check.sum
	go mod tidy -modfile=go.check.mod
	diff -u go.mod go.check.mod
	diff -u go.sum go.check.sum
	rm go.check.mod go.check.sum
.PHONY: check-tidy

build-docker-dev: ## build development image from Dockerfile.dev
 ifndef INSIDE_DEV_CONTAINER
	@echo "+ $@"
	DOCKER_BUILDKIT=1 docker build --tag pairedbrackets:dev - < Dockerfile.dev
 endif
.PHONY: build-docker-dev

ifndef INSIDE_DEV_CONTAINER
  RUN_IN_DOCKER = docker run --rm                                                                \
                             -it                                                                 \
                             -w /app                                                             \
                             --mount type=bind,consistency=delegated,source="`pwd`",target=/app  \
                             pairedbrackets:dev
endif
