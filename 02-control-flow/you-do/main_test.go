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
