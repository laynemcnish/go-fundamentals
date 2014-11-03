package main

import "fmt"

func main() {
	// create a struct called test which contains input1 with a type integer, input2 with a type string,
	// and expected with a type bool.
	type test struct {
		input1   int
		input2   string
		expected bool
	}
	/* create a slice called tests composed of 3 test structs, with
	   the first containing values 1, "/doc/", and true
	   the second containing 10, "/rmp/", and false
	   the third containing 2, "/doc/", and false */
	tests := []test{
		test{input1: 1, input2: "/doc/", expected: true},
		test{input1: 10, input2: "/rmp/", expected: false},
		test{input1: 2, input2: "/doc/", expected: false},
	}

	// print tests
	fmt.Println(tests)
}
