package main

// create a function simpleFunc which takes no arguments and returns the string "Hello Function!"
func simpleFunc() string {
	return "Hello Function!"
}

// create a function splitNum which takes an integer and returns that number divided by 2
func splitNum(i float64) float64 {
	return i / 2
}

// create a function howExciting which takes two arguments, an integer and then a string. Return a string
// which is the second argument plus as many exclamation points as the first argument
func howExciting(n int, str string) string {
	for i := 0; i < n; i++ {
		str = str + "!"
	}
	return str
}
