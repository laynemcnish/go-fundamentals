package main

// create a function inlineAssign which takes in two integers. if the double of the first integer is greater than
// the second integer, return 100, otherwise return the two input integers added together.
// !!! declare and assign a variable AFTER the 'if' but before the evaluation in the if statement !!!
func inlineAssign(i, testNum int) int {
	if double := i * 2; double > testNum {
		return 100
	}
	return testNum + i
}

// create a function evalFifty which takes one integer i and calls inlineAssign(i, 50), if the return
// value from inlineAssign is greater than or equal to 100, return "pass", otherwise "fail"
// !!! declare and assign a variable AFTER the 'if' but before the evaluation in the if statement !!!
func evalFifty(i int) string {
	if eval := inlineAssign(i, 50); eval >= 100 {
		return "pass"
	}
	return "fail"
}
