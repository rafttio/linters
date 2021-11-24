package main

import (
	"golang.org/x/tools/go/analysis/multichecker"
	"raftt.io/linters/pkg/analyzers"
)

func main() {
	multichecker.Main(analyzers.Analyzers...)
}
