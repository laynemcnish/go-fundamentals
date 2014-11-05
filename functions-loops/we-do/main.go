package main

import "strconv"

/*
  Resources: https://code.google.com/p/go-wiki/wiki/Range
             http://golang.org/pkg/strconv/

  TODO: create a function frontRange which takes slice of strings. It should return a single string which
  contains a numbered list of the strings such that ["one", "two"] would yield "1. One\n2.Two\n"
  HINT: import "strconv" and use Itoa */
func frontRange(arr []string) string {
	var finalStr string
	for i, str := range arr {
		finalStr = finalStr + strconv.Itoa(i+1) + ". " + str + "\n"
	}
	return finalStr
}

/* TODO: create a function twoReturn which takes three floats and returns two floats- the first return value should
result from multiplying the first argument by the third, and the second return value should result from
dividing the second argument by the third. */
func twoReturn(i1, i2, i3 float64) (float64, float64) {
	return i1 * i3, i2 / i3
}

/* TODO: create a variadic function coolElipsis which takes a variadic number of integers. Return the sum of each
integer after you multiply it by its index + 1. */
func coolElipsis(numbers ...int) int {
	var total int
	for i, num := range numbers {
		total += num * (i + 1)
	}
	return total
}
