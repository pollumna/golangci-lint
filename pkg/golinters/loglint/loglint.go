package loglint

import (
	"github.com/golangci/golangci-lint/v2/pkg/goanalysis"
	"github.com/pollumna/loglint/analyzer"
)

func New() *goanalysis.Linter {
	return goanalysis.
		NewLinterFromAnalyzer(analyzer.Analyzer).
		WithDesc("Check log messages according to style rules").
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}
