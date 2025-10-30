# update together with .github/workflows/ci.yml
FROM golang:1.25.3 AS go

# update together with .github/workflows/ci.yml
FROM golangci/golangci-lint:v2.6.0 AS linter

FROM go AS dev
ENV INSIDE_DEV_CONTAINER=1
WORKDIR /app
COPY --from=linter /usr/bin/golangci-lint /usr/bin/
