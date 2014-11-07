package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestCorruption(t *testing.T) {
	type test struct {
		n        int
		expected int
	}
	tests := []test{
		test{n: 4, expected: 16},
		test{n: 5, expected: 20},
	}
	for _, tt := range tests {
		result := Corruption(tt.n)
		if result != tt.expected {
			t.Errorf("corruption(%v) = %v, want %v", tt.n, result, tt.expected)
		}
	}
}

func TestFoldIt(t *testing.T) {
	type test struct {
		arr       []int
		expected1 int
		expected2 int
	}
	tests := []test{
		test{arr: []int{1, 2, 3, 4}, expected1: 10, expected2: 24},
	}
	for _, tt := range tests {
		result1, result2 := FoldIt(tt.arr)
		if result1 != tt.expected1 || result2 != tt.expected2 {
			t.Errorf("foldIt(%v) = %v, %v; wanted %v, %v", tt.arr, result1, result2, tt.expected1, tt.expected2)
		}
	}
}

func TestConcatMap(t *testing.T) {
	type test struct {
		mp       map[string][]string
		input    string
		expected string
	}
	tests := []test{
		test{mp: map[string][]string{"foo": []string{"doc", "sets", "query"}, "bar": []string{"doc", "urn", "mvc"}},
			input:    "foo",
			expected: "/doc/sets/query/"},
		test{mp: map[string][]string{"foo": []string{"doc", "sets", "query"}, "bar": []string{"doc", "urn", "mvc"}},
			input:    "bar",
			expected: "/doc/urn/mvc/"},
	}
	for _, tt := range tests {
		result := ConcatMap(tt.mp, tt.input)
		if result != tt.expected {
			t.Errorf("concatMap(%v, %v) = %v, want %v", tt.mp, tt.input, result, tt.expected)
		}
	}
}

func ExampleCorruption() {
	fmt.Println(Corruption(2))
	// Output:
	// 4 damage taken from Corruption!
	// 4 damage taken from Corruption!
	// 8
}

func TestAttackAd(t *testing.T) {
	rand.Seed(1) // ensures deterministic behavior for this test
	attacks := []string{"[insert name here] is the worst politician",
		"I find [insert name here] to be the least goodest",
		"You'll never find [insert name here] doing the right thing"}
	first := AttackAd("Ted", attacks)
	for i := 0; i < 100; i++ {
		x := AttackAd("Ted", attacks)
		if x != first {
			return
		}
	}
	t.Errorf("Only saw one string, not random.")
}
