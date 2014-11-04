package main

import "testing"

func Test_simpleFunc(t *testing.T) {
	result := simpleFunc()
	if result != "Hello Function!" {
		t.Errorf("simpleFunc() = %v, want %v", result, "Hello Function!")
	}
}

func Test_splitNum(t *testing.T) {
	type test struct {
		i        float64
		expected float64
	}
	tests := []test{
		test{i: 4, expected: 2},
		test{i: 5, expected: 2.5},
		test{i: 10.8, expected: 5.4},
	}
	for _, tt := range tests {
		result := splitNum(tt.i)
		if result != tt.expected {
			t.Errorf("splitNum(%v) = %v, want %v", tt.i, result, tt.expected)
		}
	}
}

func Test_howExciting(t *testing.T) {
	type test struct {
		n        int
		s        string
		expected string
	}
	tests := []test{
		test{n: 2, s: "golang", expected: "golang!!"},
		test{n: 10, s: "!", expected: "!!!!!!!!!!!"},
		test{n: 3, s: "_/", expected: "_/!!!"},
	}
	for _, tt := range tests {
		result := howExciting(tt.n, tt.s)
		if result != tt.expected {
			t.Errorf("howExciting(%v) = %v, want %v", tt.n, result, tt.expected)
		}
	}
}
