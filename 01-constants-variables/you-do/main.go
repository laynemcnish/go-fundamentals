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

// on only one line declare two variables x and y, both float32 type, and give them values -1, -3 respectively
var x, y float32 = -1, -3

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

// create a new struct called fishMeasure which contains a fish species' common name, scientific name,
// standard length, and weight. Use variable names- CommonName, ScienceName, Length, Weight
// length in cm should be an int and weight a float32 in kg
type fishMeasure struct {
	CommonName  string
	ScienceName string
	Length      int
	Weight      float32
}

// create a series of 3 iota constants, but skip the first iota value 0
// Name them Caesar, Crassus, and Pompeius in that order (I better not see any numbers!)
const (
	_ int = iota
	Caesar
	Crassus
	Pompeius
)

// ultra-iota: use iota to create a complex set of values for six nutrient constants.
// !!! Never type a number greater than 2 and no more than two numbers written !!!
// declare the constants in the following order:
// minerals come first and should have a value of 1
// fat is next with a value of 0
// vitamins and protein next, both with values of 2
// carbohydrates after that with a value of 3
// finally comes water with a value of 6
const (
	minerals, fat int = iota + 1, iota * 2
	vitamins, protein
	carbohydrates, _
	_, water
)
