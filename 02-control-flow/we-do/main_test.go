package main

import "testing"

func Test_FruitOrVegatable(t *testing.T) {
	type test struct {
		value    string
		expected string
	}

	tests := []test{
		test{value: "Apple", expected: "Fruit"},
		test{value: "Foo", expected: "Fruit"},
	}

	for _, tt := range tests {
		result := FruitOrVegatable(tt.value)
		if result != tt.expected {
			t.Errorf("test %q failed.  got: %q, expected: %q", tt.value, result, tt.expected)
		}

	}

}
