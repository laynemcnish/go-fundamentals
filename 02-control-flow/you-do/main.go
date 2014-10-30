package main

// create a function which takes two strings and returns true if they match, false if they don't
func wordMatch(w1, w2 string) bool {
	if w1 == w2 {
		return true
	}
	return false
}

// create a function doubleUp which takes two integers and returns the string "double up!" if one number
// is double the size of the other. Otherwise return "no double..."
func doubleUp(a, b int) string {
	if a*2 == b || a == b*2 {
		return "double up!"
	}
	return "no double..."
}

// create a function runeToNum which takes a rune and returns true if the rune's value is greater than 50, otherwise false
func runeToNum(r rune) bool {
	if r > 50 {
		return true
	}
	return false
}
