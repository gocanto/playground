package main

import (
	"fmt"

	"github.com/playground/intervals"
)

func main() {
	// Test case from the example
	input := [][]int{{1, 3}, {2, 4}, {6, 8}, {7, 9}}

	fmt.Printf("Input:  %v\n", input)

	merged := intervals.Merge(input)

	fmt.Printf("Output: %v\n", merged)
}
