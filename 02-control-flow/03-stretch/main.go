package main

/*
	References: http://blog.golang.org/laws-of-reflection
							http://golang.org/pkg/reflect/

	TODO:
	Write a function whatAmI which takes an interface and unpacks that empty interface to recover the type
	information. Switch on the value of that type such that you return "string" for a string type, "integer" for
	an int type, and "float" for a float type (32 or 64 bits). Otherwise return "unknown".
	HINT: check out the tests for the constants-variables lessons to see the reflect package in action.
*/
