package main

import "testing"

func Test_wordMatch(t *testing.T) {
	type test struct {
		name          string
		first, second string
		expected      bool
	}
	tests := []test{
		test{name: "Same words", first: "dog", second: "dog", expected: true},
	}
	for _, tt := range tests {
		result := wordMatch(tt.first, tt.second)
		if result != tt.expected {
			t.Errorf("Test %q failed. got: %d, expected %d", tt.name, result, tt.expected)
		}
	}
}

func Test_doubleUp(t *testing.T) {
	type test struct {
		name     string
		a, b     int
		expected string
	}
	tests := []test{
		test{name: "a double of b", a: 4, b: 2, expected: "double up!"},
		test{name: "a equal to b", a: 4, b: 4, expected: "no double..."},
    test{name: "a half of b", a: 4, b: 8, expected: "double up!"},
	}
	for _, tt := range tests {
		result := doubleUp(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("Test %q failed. got :%d, expected: %d", tt.name, result, tt.expected)
		}
	}
}