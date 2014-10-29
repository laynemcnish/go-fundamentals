package main

import "testing"

func Test_nutrient_values(t *testing.T) {
	type test struct {
		name     string
		value    int
		expected int
	}
	var tests = []test{
		test{name: "minerals", value: minerals, expected: 1},
		test{name: "fat", value: fat, expected: 0},
		test{name: "vitamins", value: vitamins, expected: 2},
		test{name: "protein", value: protein, expected: 2},
		test{name: "carbohydrates", value: carbohydrates, expected: 3},
		test{name: "water", value: water, expected: 6},
	}
	for _, tt := range tests {
		if tt.value != tt.expected {

			t.Errorf("Something went wrong with %s iota: expected %d, got %d", tt.name, tt.expected, tt.value)
		}
	}
}
