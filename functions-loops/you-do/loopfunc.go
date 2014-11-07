package loopfunc

import "fmt"

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
		finalString = finalString + (str + "/")
	}
	return finalString
}

/* write a function AttackAd */
