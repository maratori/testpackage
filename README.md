# testpackage <br> [![Build Status](https://travis-ci.com/maratori/testpackage.svg?branch=master)](https://travis-ci.com/maratori/testpackage) [![codecov](https://codecov.io/gh/maratori/testpackage/branch/master/graph/badge.svg)](https://codecov.io/gh/maratori/testpackage) [![codebeat badge](https://codebeat.co/badges/9c74d700-ebf8-4b76-8405-1950874576c4)](https://codebeat.co/projects/github-com-maratori-testpackage-master) [![Maintainability](https://api.codeclimate.com/v1/badges/bf753d7560c8e4aa5cf0/maintainability)](https://codeclimate.com/github/maratori/testpackage/maintainability) [![Go Report Card](https://goreportcard.com/badge/github.com/maratori/testpackage)](https://goreportcard.com/report/github.com/maratori/testpackage) [![GitHub](https://img.shields.io/github/license/maratori/testpackage.svg)](LICENSE) [![GoDoc](https://godoc.org/github.com/maratori/testpackage?status.svg)](http://godoc.org/github.com/maratori/testpackage)

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

The best way is to use [golangci-lint](https://github.com/golangci/golangci-lint).
It includes **testpackage** linter started from v1.25.0 and higher.

### Install
See [install section](https://github.com/golangci/golangci-lint#install) of readme.

### Configuration
**testpackage** is disabled by default. To enable it, add the following to your `.golangci.yml`:
```yaml
linters:
  enable:
    testpackage
```

You can also change regexp that is used to ignore files by the linter. Here is the default value.
```yaml
linters-settings:
  testpackage:
    skip-regexp: (export|internal)_test\.go
```

### Run
```shell script
golangci-lint run
```


## Usage as standalone linter

### Install
```shell script
go get -u github.com/maratori/testpackage
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

Flags:  -V      print version and exit
  -skip-regexp string
        regexp pattern to skip file by name. To not skip files use -skip-regexp="^$" (default "(export|internal)_test\\.go")
  -json
        emit JSON output
  -c int
        display offending line with this many lines of context (default -1)
  -cpuprofile string
        write CPU profile to this file
  -memprofile string
        write memory profile to this file
```


## Changelog

### [v1.0.1] - 2020-04-22

#### Changed
* No changes in linter behavior
* Use latest go version on travis-ci
* Update Makefile
* Update golangci-lint

### [v1.0.0] - 2019-11-10

#### Added
* Go Analyzer to check the name of test package
* main.go to run the analyzer
* MIT [license](LICENSE)
