package main

import "testing"

type test struct {
	name     string
	value    int
	expected int
}

func Test_triumverate_values(t *testing.T) {
	var tests = []test{
		test{name: "Ceasar", value: Caesar, expected: 1},
		test{name: "Crassus", value: Crassus, expected: 2},
		test{name: "Pompeius", value: Pompeius, expected: 3},
	}
	for _, tt := range tests {
		if tt.value != tt.expected {

			t.Errorf("Something went wrong with %s iota: expected %d, got: %d", tt.name, tt.expected, tt.value)
		}
	}
}
