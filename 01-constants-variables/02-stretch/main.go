package main

// create a new struct called fishMeasure which contains a fish species' common name, scientific name,
// standard length, and weight. Use variable names- CommonName, ScienceName, Length, Weight
// length in cm should be an int and weight a float32 in kg
type fishMeasure struct {
	CommonName  string
	ScienceName string
	Length      int
	Weight      float32
}

// create a variable called fish of fishMeasure type. Set CommonName to "cutbow",
// set ScienceName to "Oncorhynchus clarki x mykiss", Length to 39, and Weight to 0.75
var fish = fishMeasure{
	CommonName:  "cutbow",
	ScienceName: "Oncorhynchus clarki x mykiss",
	Length:      39,
	Weight:      0.75,
}
