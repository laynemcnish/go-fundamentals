package main

import "fmt"

// create a map grades with the keys "danger", "normal", "good" pointing to slices of integers.
// "danger" array should be filled with integers 0-6, "good" should be 7, 8, and 9, and
// "good" should be an array containing 10. Print each value of grades with a description

func main() {
	grades := make(map[string][]int)
	grades["danger"] = []int{0, 1, 2, 3, 4, 5, 6}
	grades["normal"] = []int{7, 8, 9}
	grades["good"] = []int{10}

	fmt.Println("Danger: ", grades["danger"])
	fmt.Println("Normal: ", grades["normal"])
	fmt.Println("Good:   ", grades["good"])
}
