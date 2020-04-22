ci: test-cover lint check-tidy
.PHONY: ci

test:
	go test -race ./...
.PHONY: test

lint:
	docker run --rm --name lint -v `pwd`:/app -w /app golangci/golangci-lint:v1.24.0 golangci-lint run
.PHONY: lint

tidy:
	go mod tidy
.PHONY: tidy

update-deps:
	go get -u -t ./...
	go mod tidy
.PHONY: update-deps

check-tidy:
	cp go.mod go.check.mod
	cp go.sum go.check.sum
	go mod tidy -modfile=go.check.mod
	diff -u go.mod go.check.mod
	diff -u go.sum go.check.sum
	rm go.check.mod go.check.sum
.PHONY: check-tidy

test-cover:
	go test -race -coverpkg ./... -coverprofile=coverage.out ./...
.PHONY: test-cover
