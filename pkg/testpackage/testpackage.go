package testpackage

import (
	"flag"
	"regexp"
	"strings"

	"golang.org/x/tools/go/analysis"
)

const (
	SkipRegexpFlagName    = "skip-regexp"
	SkipRegexpFlagUsage   = `regexp pattern to skip file by name. To not skip files use -skip-regexp="^$"`
	SkipRegexpFlagDefault = `(export|internal)_test\.go`
)

const (
	AllowPackagesFlagName    = "allow-packages"
	AllowPackagesFlagUsage   = `comma separated list of packages that don't end with _test that tests are allowed to be in`
	AllowPackagesFlagDefault = `main`
)

// NewAnalyzer returns Analyzer that makes you use a separate _test package.
func NewAnalyzer() *analysis.Analyzer {
	var (
		skipFileRegexp = SkipRegexpFlagDefault
		allowPackagesStr = AllowPackagesFlagDefault
		fs             flag.FlagSet
	)

	fs.StringVar(&skipFileRegexp, SkipRegexpFlagName, skipFileRegexp, SkipRegexpFlagUsage)
	fs.StringVar(&allowPackagesStr, AllowPackagesFlagName, allowPackagesStr, AllowPackagesFlagUsage)

	return &analysis.Analyzer{
		Name:  "testpackage",
		Doc:   "linter that makes you use a separate _test package",
		Flags: fs,
		Run: func(pass *analysis.Pass) (interface{}, error) {
			allowedPackages := strings.Split(allowPackagesStr, ",")
			skipFile, err := regexp.Compile(skipFileRegexp)
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

					allowedPackage := false
					for _, p := range allowedPackages {
						if p == packageName {
							allowedPackage = true
						}
					}

					if allowedPackage {
						continue
					}

					if !strings.HasSuffix(packageName, "_test") {
						pass.Reportf(f.Name.Pos(), "package should be `%s_test` instead of `%s`", packageName, packageName)
					}
				}
			}

			return nil, nil
		},
	}
}
