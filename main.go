package main

import (
	"flag"

	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/maratori/testpackage/pkg/testpackage"
)

func main() {
	skip := flag.String("skip-regexp", testpackage.DefaultSkipRegexp, `regexp pattern to skip file by name. To not skip files use -skip-regexp="^$"`) // nolint:lll
	singlechecker.Main(testpackage.NewAnalyzer(skip))
}
