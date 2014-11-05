package main

import "testing"

func Test_frontRange(t *testing.T) {
	type test struct {
		arr      []string
		expected string
	}
	tests := []test{
		test{arr: []string{"One", "Two"}, expected: "1. One\n2. Two\n"},
		test{arr: []string{"Thing", "Other", "Another"}, expected: "1. Thing\n2. Other\n3. Another\n"},
	}
	for _, tt := range tests {
		result := frontRange(tt.arr)
		if result != tt.expected {
			t.Errorf("frontRange(%v) = \n%v, want \n%v", tt.arr, result, tt.expected)
		}
	}
}

func Test_twoReturn(t *testing.T) {
	type test struct {
		i1        float64
		i2        float64
		i3        float64
		expected1 float64
		expected2 float64
	}
	tests := []test{
		test{i1: 10, i2: 20, i3: 5, expected1: 50, expected2: 4},
	}
	for _, tt := range tests {
		result1, result2 := twoReturn(tt.i1, tt.i2, tt.i3)
		if result1 != tt.expected1 || result2 != tt.expected2 {
			t.Errorf("twoReturn(%v, %v, %v) = %v,%v; want %v, %v", tt.i1, tt.i2, tt.i3, result1, result2, tt.expected1, tt.expected2)
		}
	}
}

func Test_coolElipsis(t *testing.T) {
	type test struct {
		nums     []int
		expected int
	}
	tests := []test{
		test{nums: []int{1, 2, 3, 4}, expected: 30},
	}
	for _, tt := range tests {
		result := coolElipsis(tt.nums...)
		if result != tt.expected {
			t.Errorf("coolElipsis(%v) = %v, want %v", tt.nums, result, tt.expected)
		}
	}
}
