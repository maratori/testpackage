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

	skip := testpackage.DefaultSkipRegexp
	analysistest.Run(t, testdata, testpackage.NewAnalyzer(&skip))
}

func TestAnalyzer_Bad(t *testing.T) {
	testdata, err := filepath.Abs("testdata/bad")
	if err != nil {
		t.FailNow()
	}

	skip := testpackage.DefaultSkipRegexp
	analysistest.Run(t, testdata, testpackage.NewAnalyzer(&skip))
}
