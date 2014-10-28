package main

import "strconv"

// declare a variable string called sandwich and give it a value of "pork chop"
var sandwich = "pork chop"

// declare four variables first, second, third, fourth, each an int with a value of 1, 2, 3, 4
// do not write 'var' or 'int' more than once!
var (
	first  int = 1
	second     = 2
	third      = 3
	fourth     = 4
)

// declare a constant typedFloat, give it a type of float32, set it equal to 4.0
const typedFloat float32 = 4.0

// declare a constant untypedFloat, do NOT assign it a type (i.e. let it be the default type)- set it equal to 4.0
const untypedFloat = 4.0

// declare a variable bigSixtyFour, assign it a type of int64 and set it to its largest possible positive number
var bigSixtyFour int64 = 9223372036854775807

// declare a variable negaNumber, assign it a type whose lowest possible negative number is -2147483648. Set that number.
var negaNumber int32 = -2147483648

// declare a variable named char with a rune type, set it to 'Q'
var char rune = 'Q'

// declare a variable letterFromChar- create it by converting char to a string
var letterFromChar = string(char)

// declare a variable numberFromChar- create it by converting char to a 64 bit integer
var numberFromChar = int64(char)

// declare a variable cryBaby, create it by adding letterFromChar and numberFromChar converted to a string
var cryBaby = letterFromChar + string(numberFromChar)

// declare a variable charIntString which is set to a string type representation of numberFromChar
// HINT: import "strconv" --- strconv.Itoa(T int32)
var charIntString = strconv.Itoa(int(numberFromChar))

// EXTRA CREDIT

// declare constants that have four integer iota values ranging from 0 to 3
// Name them Agape, Phileo, Storge and Eros (if you've written any numbers, you've done it wrong!)
const (
	Agape int = iota
	Phileo
	Storge
	Eros
)

// create a new struct called fishMeasure which contains a fish species' common name, scientific name,
// standard length, and weight. Use variable names- CommonName, ScienceName, Length, Weight
// length in cm should be an int and weight a float32 in kg
type fishMeasure struct {
	CommonName  string
	ScienceName string
	Length      int
	Weight      float32
}
