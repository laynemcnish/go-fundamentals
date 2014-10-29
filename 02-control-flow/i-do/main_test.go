package mathUtil

import (
	"testing"
)

func Test_Max(t *testing.T) {

	type test struct {
		name     string
		x, y     int
		expected int
	}

	tests := []test{
		test{name: "x > y", x: 10, y: 5, expected: 10},
		test{name: "y > x", x: 5, y: 10, expected: 10},
		test{name: "y == x", x: 5, y: 5, expected: 5},
		test{name: "-x < +y", x: -5, y: 5, expected: 5},
	}

	for _, tt := range tests {
		result := Max(tt.x, tt.y)
		if result != tt.expected {
			t.Errorf("Test %q failed.  got: %d, expected: %d", tt.name, result, tt.expected)
		}
	}

}
