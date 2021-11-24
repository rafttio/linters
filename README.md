# Raftt
[![CI](https://github.com/rafttio/linters/actions/workflows/go-test.yml/badge.svg)](https://github.com/rafttio/linters/actions/workflows/go-test.yml)
# Raftt linters

This package contains Go linters for use in Raftt projects.

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
