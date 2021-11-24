//go:build dontbuildforgitlabci
// +build dontbuildforgitlabci

// this package is needed to make go always include golangci-lint in the dependencies
// so that we are building the plugin against their correct dep versions

package build

import (
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
)
