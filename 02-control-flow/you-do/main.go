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

// Knowing if it is a rune is nice, but let us see if the rune is a letter in the alphabet. Write a function
// called runeInAlphabet which takes a rune and returns "capital letter" if the letter is in the English alphabet
// and "lowercase letter" if it is lowercase, and "not a letter" if it is not in the alphabet.
// HINT: Ken Thompson made this self-synchronizing design outline with Rob Pike on Sept 2, 1992 on a placemat in a New Jersey diner.
func runeInAlphabet(r rune) string {
	switch {
	case r >= 65 && r <= 90:
		return "capital letter"
	case r >= 97 && r <= 122:
		return "lowercase letter"
	default:
		return "not a letter"
	}
}
