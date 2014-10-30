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
			t.Errorf("Test %q failed. got: %d, expected %v", tt.name, result, tt.expected)
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
		test{name: "a not double, not equal to b", a: 1, b: 7, expected: "no double..."},
		test{name: "a half of b", a: 4, b: 8, expected: "double up!"},
	}
	for _, tt := range tests {
		result := doubleUp(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("Test %q failed. got :%d, expected: %v", tt.name, result, tt.expected)
		}
	}
}

func Test_runeToNum(t *testing.T) {
	type test struct {
		r        rune
		expected bool
	}
	tests := []test{
		test{r: 'F', expected: true},
		test{r: '/', expected: false},
	}
	for _, tt := range tests {
		result := runeToNum(tt.r)
		if result != tt.expected {
			t.Errorf("runeToNum(%d) = %d, want %v", tt.r, result, tt.expected)
		}
	}
}

func Test_multipleByte(t *testing.T) {
	type test struct {
		i        int
		expected string
	}
	tests := []test{
		test{i: 1024, expected: "kibibyte"},
		test{i: 1025, expected: "inexact"},
		test{i: 1048576, expected: "mebibyte"},
		test{i: 1073741824, expected: "gibibyte"},
	}
	for _, tt := range tests {
		result := multipleByte(tt.i)
		if result != tt.expected {
			t.Errorf("multipleByte(%d) = %d, want %v", tt.i, result, tt.expected)
		}
	}
}

func Test_correctByte(t *testing.T) {
	type test struct {
		i        int
		expected int
	}
	tests := []test{
		test{i: 1000, expected: 1024},
		test{i: 1000000, expected: 1048576},
		test{i: 1000000000, expected: 1073741824},
		test{i: 5, expected: 5},
	}
	for _, tt := range tests {
		result := correctByte(tt.i)
		if result != tt.expected {
			t.Errorf("correctByte(%d) = %d, want %v", tt.i, result, tt.expected)
		}
	}
}
