package main

import (
	"fmt"
	"math/rand"
	"strings"
)

// write a function Corruption which takes in an integer. Print out "4 damage taken from Corruption!" as many times
// as the integer argument. Return the total damage taken at the end.

func Corruption(n int) int {
	for i := 0; i < n; i++ {
		fmt.Println("4 damage taken from Corruption!")
	}
	return n * 4
}

// write a function FoldIt which takes a slice of integers and returns two integers. The first return
// should be the result of adding all the integers in the slice, the second a result from multiplying.

func FoldIt(arr []int) (int, int) {
	var a int
	b := 1
	for _, num := range arr {
		a += num
		b *= num
	}
	return a, b
}

/* write a function ConcatMap with two arguments; a map and a string type. The map should contain string keys
which point to string slices. Your function should use its string argument to select the string key from the map
and concat all the strings in the slice value. Place "/" between each string. Return the concatenated
string with '/' as the first character. */

func ConcatMap(mp map[string][]string, input string) string {
	finalString := "/"
	for _, str := range mp[input] {
		finalString += str + "/"
	}
	return finalString
}

/* write a function AttackAd which takes a string and a slice of strings. Each string in the slice is expected
to contain "[insert name here]". Pick a string from the slice at random and replace "[insert name here]" with
the first string argument and return it. Each instance of "[insert name here]" should be replaced
HINT: import "math/rand" and "strings", they have the tools you need (Replace and Intn) */

func AttackAd(name string, attacks []string) string {
	anonAttack := attacks[rand.Intn(len(attacks))]
	attack := strings.Replace(anonAttack, "[insert name here]", name, -1)
	return attack
}
