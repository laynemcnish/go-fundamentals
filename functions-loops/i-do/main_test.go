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
