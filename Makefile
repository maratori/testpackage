.PHONY: ci
ci: testcov lint istidy

.PHONY: test
test:
	go test -race ./...

.PHONY: lint
lint:
	docker run --rm --name lint -v `pwd`:/app -w /app golangci/golangci-lint:v1.24.0 golangci-lint run

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: testcov
testcov:
	go test -race -coverpkg ./... -coverprofile=coverage.out ./...

.PHONY: istidy
istidy:
	go mod tidy
	if [[ `git status --porcelain go.mod` ]]; then git diff -- go.mod ; echo "go.mod is outdated, please run go mod tidy" ; exit 1; fi
	if [[ `git status --porcelain go.sum` ]]; then git diff -- go.sum ; echo "go.sum is outdated, please run go mod tidy" ; exit 1; fi
