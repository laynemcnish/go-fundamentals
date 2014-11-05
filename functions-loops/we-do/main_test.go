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
