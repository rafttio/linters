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

func bazz() (int, float64, error) {
	return 0, 1.1, nil
}

func Tests() {
	foo()     // ok
	bar()     // want "call discards return value"
	_ = bar() // ok

	a := bar() // ok
	a()

	defer bar()   // want "call discards return value"
	defer bar()() // ok

	bazz()           // want "call discards return value"
	_, _, _ = bazz() //ok
}
