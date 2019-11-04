/*
testpackage is golang linter that makes you use a separate `_test` package.

Motivation

According to blackbox testing approach, you should not use unexported functions and methods from source code in tests.

Go allows to place tests in a separate package with suffix `_test`.
For example, tests for `store` package can be in the same package or in the package `store_test`.
In the second case, you have to import the source code into tests so only exported things are available.

The linter reports if a test is in a package without suffix `_test`.
If you really need to test unexported function, then put the test into file `XXX_internal_test.go`.
The linter skips such files by default.
It also skips the file `export_test.go` by default (see the last article below).

More detailed articles on this topic:
 * Next level Go testing by Tit Petric https://scene-si.org/2019/04/15/next-level-go-testing#public-vs-private-tests-apis
 * 5 simple tips and tricks for writing unit tests in #golang by Mat Ryer https://medium.com/@matryer/5-simple-tips-and-tricks-for-writing-unit-tests-in-golang-619653f90742
 * 5 advanced testing techniques in Go by Alan Braithwaite https://segment.com/blog/5-advanced-testing-techniques-in-go/#use-a-separate-test-package
 * Golang Trick: Export unexport method for test by lysu https://medium.com/@robiplus/golang-trick-export-for-test-aa16cbd7b8cd
*/
package main
