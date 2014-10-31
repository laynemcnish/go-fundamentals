package main

/*
	TODO:
	Let us see if a rune is a letter in the alphabet. Write a function
	called runeInAlphabet which takes a rune and returns "capital letter" if the letter is in the English alphabet
	and "lowercase letter" if it is lowercase, and "not a letter" if it is not in the alphabet.
	HINT: Ken Thompson made this self-synchronizing design outline with Rob Pike on Sept 2, 1992 on a placemat in a New Jersey diner.
*/
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
