package main

import "reflect"

func whatAmI(v interface{}) {
	switch t := reflect.TypeOf(v); t {
	case t.Kind() == reflect.String:
		return "String"

	default:
		return "Unknown"
	}

}
