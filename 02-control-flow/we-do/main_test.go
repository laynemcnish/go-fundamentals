package main

import "testing"

func Test_FruitOrVegetable(t *testing.T) {
	type test struct {
		value    string
		expected string
	}

	tests := []test{
		test{value: "Apple", expected: "Fruit"},
		test{value: "Orange", expected: "Fruit"},
		test{value: "Foo", expected: "Neither"},
		test{value: "Onion", expected: "Vegetable"},
		test{value: "Celery", expected: "Vegetable"},
	}

	for _, tt := range tests {
		result := FruitOrVegetable(tt.value)
		if result != tt.expected {
			t.Errorf("test %q failed.  got: %q, expected: %q", tt.value, result, tt.expected)
		}

	}

}
