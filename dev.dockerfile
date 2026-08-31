# update together with .github/workflows/ci.yml and .github/latest-deps/go.mod (for minor version change)
FROM golang:1.27.0 AS go

# update together with .github/workflows/ci.yml
FROM golangci/golangci-lint:v2.13.2 AS linter

FROM go AS dev
ENV INSIDE_DEV_CONTAINER=1
WORKDIR /app
COPY --from=linter /usr/bin/golangci-lint /usr/bin/
