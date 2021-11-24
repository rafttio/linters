
# Raftt linters

This package contains several Go linters for use in Raftt projects.

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
