package main

/*
	Reference: https://golang.org/ref/spec#Iota

	TODO:
	Declare constants that have four integer iota values ranging from 0 to 3
	Name them Agape, Phileo, Storge and Eros

	Hint: if you've written any numbers, you've done it wrong!
*/

const (
	Agape int = iota
	Phileo
	Storge
	Eros
)
