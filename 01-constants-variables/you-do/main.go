package main

import "strconv"

// declare a constant typedFloat, give it a type of float32, set it equal to 4.0
const typedFloat float32 = 4.0

// declare a constant untypedFloat, do NOT assign it a type (i.e. let it be the default type)- set it equal to 4.0
const untypedFloat = 4.0

// declare a variable bigSixtyFour, assign it a type of int64 and set it to its largest possible positive number
var bigSixtyFour int64 = 9223372036854775807

// declare a variable negaNumber, assign it a type whose lowest possible negative number is -2147483648. Set that number.
var negaNumber int32 = -2147483648

// declare a variable named char with a rune type, set it to 'H'
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
