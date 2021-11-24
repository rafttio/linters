package tests

type Returned struct{}

func foo() {
	println("foo called")
}

func bar() func() {
	return func() {
		println("called closure")
	}
}

func bazz() (int, struct{}, error) {
	return 0, struct{}{}, nil
}

func Tests() {
	foo()     // ok
	bar()     // want "call discards return value"
	_ = bar() // ok

	a := bar() // ok
	a()

	defer bar()   // want "call discards return value"
	defer bar()() // ok

	//nolint:errcheck
	bazz()           // want "call discards return value"
	_, _, _ = bazz() //ok
}
