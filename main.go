package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/maratori/testpackage/pkg/testpackage"
)

func main() {
	singlechecker.Main(testpackage.NewAnalyzer())
}
