package testpackage_test

import (
	"path/filepath"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/maratori/testpackage/pkg/testpackage"
)

func TestAnalyzer_Good(t *testing.T) {
	testdata, err := filepath.Abs("testdata/good")
	if err != nil {
		t.FailNow()
	}

	analysistest.Run(t, testdata, testpackage.NewAnalyzer())
}

func TestAnalyzer_Bad(t *testing.T) {
	testdata, err := filepath.Abs("testdata/bad")
	if err != nil {
		t.FailNow()
	}

	analysistest.Run(t, testdata, testpackage.NewAnalyzer())
}

func TestAnalyzer_InvalidRegexp(t *testing.T) {
	invalid := `\Ca`
	analyzer := testpackage.NewAnalyzer()
	err := analyzer.Flags.Set(testpackage.SkipRegexpFlagName, invalid)

	if err != nil {
		t.FailNow()
	}

	result, err := analyzer.Run(nil)

	if err == nil {
		t.FailNow()
	}

	if err.Error() != "error parsing regexp: invalid escape sequence: `\\C`" {
		t.Fatalf("Wrong error %q", err.Error())
	}

	if result != nil {
		t.FailNow()
	}
}
