package main

import "fmt"

/*
  Resources: https://golang.org/doc/effective_go.html#for

  TODO: create a 2D slice called threeByThree which contains a length 3 slice with a 3 length slice of integers
at each index. Starting at 0, increase each grid point on the matrix by 1 such that printing threeByThree
will yield [[0,1,2] [3,4,5] [6,7,8]]. You can set them manually, but try to use two for loops.
Print each index of threeByThree. */

func main() {
	threeByThree := make([][]int, 3)
	// manually creating
	// threeByThree[0] = make([]int, 3)
	// threeByThree[1] = make([]int, 3)
	// threeByThree[2] = make([]int, 3)
	//
	// threeByThree[0][0] = 0
	// threeByThree[0][1] = 1
	// threeByThree[0][2] = 2
	// threeByThree[1][0] = 3
	// threeByThree[1][1] = 4
	// threeByThree[1][2] = 5
	// threeByThree[2][0] = 6
	// threeByThree[2][1] = 7
	// threeByThree[2][2] = 8

	// using for loops to create
	var count int
	for i := 0; i < 3; i++ {
		threeByThree[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			threeByThree[i][j] = count
			count = count + 1
		}
	}
	fmt.Println(threeByThree)
}
