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

// write a function multipleByte which returns "kibibyte" if an integer matches one kebibyte value,
// "mebibyte" if it matches the value of one mebibyte, and "gibibyte" if it matches the value of one gibibyte.
// otherwise it should return "inexact".
func multipleByte(i int) string {
	switch i {
	case 1024:
		return "kibibyte"
	case 1048576:
		return "mebibyte"
	case 1073741824:
		return "gibibyte"
	default:
		return "inexact"
	}
}

// Similar to the previous function, let us write a function correctByte which will return the binary value
// of a metric decimal value (e.g. 1000 will return 1024). I only care about exact matches of 1000, 1000**2, and
// 1000**3; everything else should just return the decimal value.
func correctByte(i int) int {
	switch i {
	case 1000:
		return 1024
	case 1000000:
		return 1048576
	case 1000000000:
		return 1073741824
	default:
		return i
	}
}

// create a function threeString which takes three strings. If the first string is "stop", return the second string.
// If the first string is "go" return the third string. If the first string is "golang" return "awesome!!!"
// otherwise return the first string.
func threeString(s1, s2, s3 string) string {
	switch s1 {
	case "stop":
		return s2
	case "go":
		return s3
	case "golang":
		return "awesome!!!"
	default:
		return s1
	}
}

// create a function topThree which takes in a boolean and an integer. If the boolean is true, switch on the integer
// such that 1 returns "1.", 2 returns "2.", 3 returns "3." and any other integer returns "not good enough".
// When the boolean is false, switch on the integer such that 1 returns "One:", 2 returns "Two:" and 3 "Three:"
// and anything else returns "not good enough".
func topThree(b bool, i int) string {
	if b == true {
		switch i {
		case 1:
			return "1."
		case 2:
			return "2."
		case 3:
			return "3."
		default:
			return "not good enough"
		}
	} else {
		switch i {
		case 1:
			return "One:"
		case 2:
			return "Two:"
		case 3:
			return "Three:"
		default:
			return "not good enough"
		}
	}
}
