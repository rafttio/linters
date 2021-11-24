package main

import (
	"golang.org/x/tools/go/analysis"
	"raftt.io/linters/pkg/analyzers"
)

type analyzerPlugin struct{}

func (*analyzerPlugin) GetAnalyzers() []*analysis.Analyzer {
	return analyzers.Analyzers
}

var AnalyzerPlugin analyzerPlugin //nolint:deadcode
