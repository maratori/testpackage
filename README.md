# testpackage <br> [![go.mod version][go-img]][go-url] [![CI][ci-img]][ci-url] [![Codecov][codecov-img]][codecov-url] [![Codebeat][codebeat-img]][codebeat-url] [![Maintainability][codeclimate-img]][codeclimate-url] [![Go Report Card][goreportcard-img]][goreportcard-url] [![License][license-img]][license-url] [![Go Reference][godoc-img]][godoc-url]

**testpackage** is a golang linter that makes you use a separate `_test` package.

## Motivation

According to blackbox testing approach, you should not use unexported functions and methods from source code in tests.

Go allows to place tests in a separate package with suffix `_test`.
For example, tests for `store` package can be in the same package or in the package `store_test`.
In the second case, you have to import the source code into tests so only exported things are available.

The linter reports if a test is in a package without suffix `_test`.
If you really need to test unexported function, then put the test into file `XXX_internal_test.go`.
The linter skips such files by default.
It also skips the file `export_test.go` by default (see the last article below).

More detailed articles on this topic:
 * [Next level Go testing](https://scene-si.org/2019/04/15/next-level-go-testing#public-vs-private-tests-apis) by Tit Petric
 * [5 simple tips and tricks for writing unit tests in #golang](https://medium.com/@matryer/5-simple-tips-and-tricks-for-writing-unit-tests-in-golang-619653f90742) by Mat Ryer
 * [5 advanced testing techniques in Go](https://segment.com/blog/5-advanced-testing-techniques-in-go/#use-a-separate-test-package) by Alan Braithwaite
 * [Golang Trick: Export unexport method for test](https://medium.com/@robiplus/golang-trick-export-for-test-aa16cbd7b8cd) by lysu

## Usage

The best way is to use [golangci-lint](https://golangci-lint.run/). It includes [testpackage](https://golangci-lint.run/usage/linters/#list-item-testpackage) and a lot of other great linters.

### Install
See [official site](https://golangci-lint.run/usage/install/).

### Configuration
**testpackage** is disabled by default. To enable it, add the following to your `.golangci.yml`:
```yaml
linters:
  enable:
    testpackage
```

You can also change the regexp that is used to ignore files by the linter,
and the list of packages that are allowed by default.

Here are the default values:
```yaml
linters-settings:
  testpackage:
    skip-regexp: (export|internal)_test\.go
    allow-packages:
      - main
```

### Run
```shell script
golangci-lint run
```

## Usage as standalone linter

### Install
```shell script
go install github.com/maratori/testpackage
```

### Run
```shell script
testpackage ./...
```
or
```shell script
testpackage -skip-regexp="^$" ./...
```

### Command line arguments
```shell script
testpackage -help
```
```
testpackage: linter that makes you use a separate _test package

Usage: testpackage [-flag] [package]


Flags:
  -skip-regexp string
        regexp pattern to skip file by name. To not skip files use -skip-regexp="^$" (default "(export|internal)_test\\.go")
  -allow-packages string
        comma separated list of packages that don't end with _test that tests are allowed to be in (default "main")
  -V    print version and exit
  -c int
        display offending line with this many lines of context (default -1)
  -cpuprofile string
        write CPU profile to this file
  -debug string
        debug flags, any subset of "fpstv"
  -fix
        apply all suggested fixes
  -flags
        print analyzer flags in JSON
  -json
        emit JSON output
  -memprofile string
        write memory profile to this file
  -test
        indicates whether test files should be analyzed, too (default true)
  -trace string
        write trace log to this file
```

## Contributors
* [maratori](https://github.com/maratori)
* [G-Rath](https://github.com/G-Rath)

## Changelog

### [v1.1.1] - 2023-03-07

#### Changed
* Update golang to 1.20
* Update dependencies
* Update golangci-lint to v1.51.2
* Update Makefile

### [v1.1.0] - 2022-06-22

#### Changed
* Allow tests in `main` package by default and add flag `-allow-packages` to allow tests without `_test` suffix (thanks [G-Rath](https://github.com/G-Rath))
* Update golang to 1.18
* Migrate to [github actions](https://github.com/maratori/testpackage/actions/workflows/ci.yaml) from travis-ci
* Update golangci-lint to v1.46.2

### [v1.0.1] - 2020-04-22

#### Changed
* No changes in linter behavior
* Use latest go version on travis-ci
* Update Makefile
* Update golangci-lint

### [v1.0.0] - 2019-11-10

#### Added
* Go Analyzer to check the name of test package
* `main.go` to run the analyzer
* MIT [license](LICENSE)


[go-img]: https://img.shields.io/github/go-mod/go-version/maratori/testpackage
[go-url]: /go.mod
[ci-img]: https://github.com/maratori/testpackage/actions/workflows/ci.yaml/badge.svg
[ci-url]: https://github.com/maratori/testpackage/actions/workflows/ci.yaml
[codecov-img]: https://codecov.io/gh/maratori/testpackage/branch/main/graph/badge.svg?token=Pa334H8xEh
[codecov-url]: https://codecov.io/gh/maratori/testpackage
[codebeat-img]: https://codebeat.co/badges/c5ab864e-dbe5-424a-93ae-75ad98c1ea55
[codebeat-url]: https://codebeat.co/projects/github-com-maratori-testpackage-main
[codeclimate-img]: https://api.codeclimate.com/v1/badges/bf753d7560c8e4aa5cf0/maintainability
[codeclimate-url]: https://codeclimate.com/github/maratori/testpackage/maintainability
[goreportcard-img]: https://goreportcard.com/badge/github.com/maratori/testpackage
[goreportcard-url]: https://goreportcard.com/report/github.com/maratori/testpackage
[license-img]: https://img.shields.io/github/license/maratori/testpackage.svg
[license-url]: /LICENSE
[godoc-img]: https://pkg.go.dev/badge/github.com/maratori/testpackage.svg
[godoc-url]: https://pkg.go.dev/github.com/maratori/testpackage
