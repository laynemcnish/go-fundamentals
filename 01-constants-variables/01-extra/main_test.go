package main

import "testing"

func Test_intIota_values(t *testing.T) {
	type test struct {
		name     string
		value    int
		expected int
	}
	var tests = []test{
		test{name: "Agape", value: Agape, expected: 0},
		test{name: "Phileo", value: Phileo, expected: 1},
		test{name: "Storge", value: Storge, expected: 2},
		test{name: "Eros", value: Eros, expected: 3},
	}

	for _, tt := range tests {
		if tt.value != tt.expected {

			t.Errorf("Something went wrong with %s iota: expected: %d, got: %d", tt.name, tt.expected, tt.value)
		}
	}
}
