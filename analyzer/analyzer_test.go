package analyzer

import (
	"path/filepath"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	td, err := filepath.Abs("../testdata")
	if err != nil {
		t.Fatal(err)
	}
	analysistest.Run(t, td, &Analyzer, "logs")
}
