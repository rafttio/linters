package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"
	"raftt.io/linters/pkg/analyzers"
)

func main() {
	singlechecker.Main(&analyzers.DiscardAnalyzer)
}
