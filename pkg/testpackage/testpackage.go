package testpackage

import (
	"regexp"
	"strings"

	"golang.org/x/tools/go/analysis"
)

const DefaultSkipRegexp = `(export|internal)_test\.go`

// NewAnalyzer returns Analyzer that make you use a separate _test package
func NewAnalyzer(skipFileRegexp *string) *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: "testpackage",
		Doc:  "linter that make you use a separate _test package",
		Run: func(pass *analysis.Pass) (interface{}, error) {
			skipFile, err := regexp.Compile(*skipFileRegexp)
			if err != nil {
				return nil, err
			}

			for _, f := range pass.Files {
				fileName := pass.Fset.Position(f.Pos()).Filename
				if skipFile.MatchString(fileName) {
					continue
				}

				if strings.HasSuffix(fileName, "_test.go") {
					packageName := f.Name.Name
					if !strings.HasSuffix(packageName, "_test") {
						pass.Reportf(f.Name.Pos(), `package should be "%s_test" instead of "%s"`, packageName, packageName)
					}
				}
			}

			return nil, nil
		},
	}
}
