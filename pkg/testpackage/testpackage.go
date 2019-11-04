package testpackage

import (
	"strings"

	"golang.org/x/tools/go/analysis"
)

// Analyzer that make you use a separate _test package
var Analyzer = &analysis.Analyzer{
	Name: "testpackage",
	Doc:  "linter that make you use a separate _test package",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		fileName := pass.Fset.Position(f.Pos()).Filename
		if strings.HasSuffix(fileName, "_test.go") {
			packageName := f.Name.Name
			if !strings.HasSuffix(packageName, "_test") {
				pass.Reportf(f.Pos(), `package should be "%s_test" instead of "%s"`, packageName, packageName)
			}
		}
	}
	return nil, nil
}
