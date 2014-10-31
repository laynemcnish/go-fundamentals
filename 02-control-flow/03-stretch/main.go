package main

import "reflect"

/*
	References: http://blog.golang.org/laws-of-reflection
							http://golang.org/pkg/reflect/

	TODO:
	Write a function whatAmI which takes an interface and unpacks that empty interface to recover the type
	information. Switch on the value of that type such that you return "string" for a string type, "integer" for
	an int type, and "float" for a float type (32 or 64 bits). Otherwise return "unknown".
	HINT: check out the tests for the constants-variables lessons to see the reflect package in action.
*/
func whatAmI(v interface{}) string {
	switch t := reflect.TypeOf(v); {
	case t.Kind() == reflect.String:
		return "string"
	case t.Kind() == reflect.Int:
		return "int"
	case t.Kind() == reflect.Float32 || t.Kind() == reflect.Float64:
		return "float"
	default:
		return "unknown"
	}
}
