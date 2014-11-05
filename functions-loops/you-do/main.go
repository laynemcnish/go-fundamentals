package main

import "fmt"

// write a function corruption which takes in an integer. Print out "4 damage taken from Corruption!" as many times
// as the integer argument. Return the total damage taken at the end.
func corruption(n int) int {
	for i := 0; i < n; i++ {
		fmt.Println("4 damage taken from Corruption!")
	}
	return n * 4
}

// write a function foldIt which takes a slice of integers and returns two integers. The first return
// should be the result of adding all the integers in the slice, the second a result from multiplying.
func foldIt(arr []int) (int, int) {
	var a int
	b := 1
	for _, num := range arr {
		a += num
		b = b * num
	}
	return a, b
}
