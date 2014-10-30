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
