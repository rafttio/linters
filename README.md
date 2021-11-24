# Raftt
[![CI](https://github.com/rafttio/linters/actions/workflows/go-test.yml/badge.svg)](https://github.com/rafttio/linters/actions/workflows/go-test.yml)
# Raftt linters

This package contains Go linters for use in Raftt projects.

# Linters
## `discardedreturn`

This linter checks for discarded return values from functions:

```go
func foo() func() {
    ...
}

 
foo()         // error: call discards return value
_ = foo()     // OK
defer foo()   // error: call discards return value
defer foo()() // OK
```

by default, the linter ignores basic types (e.g integers, bools) and errors.

# Running
You can build the golangci-lint plugin by running `go build -buildmode=plugin golangci-lint/plugin.go`.

You can use the supplied Dockerfile to build and run the linter, or

`go run cmd/all/main.go`

