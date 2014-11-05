package main

import "testing"

func Test_corruption(t *testing.T) {
	type test struct {
		n        int
		expected int
	}
	tests := []test{
		test{n: 4, expected: 16},
		test{n: 5, expected: 20},
	}
	for _, tt := range tests {
		result := corruption(tt.n)
		if result != tt.expected {
			t.Errorf("corruption(%v) = %v, want %v", tt.n, result, tt.expected)
		}
	}
}

func Test_foldIt(t *testing.T) {
	type test struct {
		arr       []int
		expected1 int
		expected2 int
	}
	tests := []test{
		test{arr: []int{1, 2, 3, 4}, expected1: 10, expected2: 24},
	}
	for _, tt := range tests {
		result1, result2 := foldIt(tt.arr)
		if result1 != tt.expected1 || result2 != tt.expected2 {
			t.Errorf("foldIt(%v) = %v, %v; wanted %v, %v", tt.arr, result1, result2, tt.expected1, tt.expected2)
		}
	}
}
