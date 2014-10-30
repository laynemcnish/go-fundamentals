package main

import "testing"

func Test_runeInAlphabet(t *testing.T) {
	type test struct {
		r        rune
		expected string
	}
	tests := []test{
		test{r: 'F', expected: "capital letter"},
		test{r: 'c', expected: "lowercase letter"},
		test{r: 'Z', expected: "capital letter"},
		test{r: 'a', expected: "lowercase letter"},
		test{r: '/', expected: "not a letter"},
	}
	for _, tt := range tests {
		result := runeInAlphabet(tt.r)
		if result != tt.expected {
			t.Errorf("runeInAlphabet(%d) = %d, want %v", tt.r, result, tt.expected)
		}
	}
}
