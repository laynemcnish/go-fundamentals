package main

import "testing"

func Test_inlineAssign(t *testing.T) {
	type test struct {
		i        int
		num      int
		expected int
	}
	tests := []test{
		test{i: 5, num: 25, expected: 30},
		test{i: 5, num: 9, expected: 100},
		test{i: 5, num: 1, expected: 100},
		test{i: 25, num: 100, expected: 125},
	}
	for _, tt := range tests {
		result := inlineAssign(tt.i, tt.num)
		if result != tt.expected {
			t.Errorf("inlineAssing(%v, %v) = %v, want %d", tt.i, tt.num, result, tt.expected)
		}
	}
}
