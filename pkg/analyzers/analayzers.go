package analyzers

import "golang.org/x/tools/go/analysis"

var Analyzers = []*analysis.Analyzer{
	&DiscardAnalyzer,
}
