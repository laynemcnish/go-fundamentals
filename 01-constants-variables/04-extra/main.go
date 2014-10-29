package main

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
