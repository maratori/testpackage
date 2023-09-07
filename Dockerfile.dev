# update together with .github/workflows/ci.yml
FROM golang:1.21.1 AS go

# update together with .github/workflows/ci.yml
FROM golangci/golangci-lint:v1.54.2 AS linter

FROM go AS dev
ENV INSIDE_DEV_CONTAINER 1
WORKDIR /app
COPY --from=linter /usr/bin/golangci-lint /usr/bin/
