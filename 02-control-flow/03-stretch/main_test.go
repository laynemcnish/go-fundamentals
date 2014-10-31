package main

import "testing"

func Test_whatAmI(t *testing.T) {
	type test struct {
		v        interface{}
		expected string
	}
	var thirtyTwo float32
	tests := []test{
		test{v: "string", expected: "string"},
		test{v: 30, expected: "int"},
		test{v: 1.5, expected: "float"},
		test{v: thirtyTwo, expected: "float"},
	}
	for _, tt := range tests {
		result := whatAmI(tt.v)
		if result != tt.expected {
			t.Errorf("whatAmI(%v) = %v, want %v", tt.v, result, tt.expected)
		}
	}
}
