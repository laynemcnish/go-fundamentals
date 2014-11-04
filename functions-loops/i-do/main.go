package main

import "fmt"

// create a function simpleFunc which takes no arguments and returns the string "Hello Function!"
func simpleFunc() string {
	return "Hello Function!"
}

// create a function splitNum which takes an integer and returns that number divided by 2
func splitNum(i float64) float64 {
	return i / 2
}

func main() {
	fmt.Println(simpleFunc())
}
