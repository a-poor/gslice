package gslice_test

import (
	"fmt"

	"github.com/a-poor/gslice"
)

func ExampleAppendAt() {
	start := []int{1, 2, 4, 5}

	// Add a single value to the middle of the slice
	result, err := gslice.AppendAt(start, 2, 3)

	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", result)
	// Output: [1 2 3 4 5]
}

func ExampleAppendAt_multiple() {
	start := []int{1, 5}

	// Add multiple elements at once
	result, err := gslice.AppendAt(start, 1, 2, 3, 4)

	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", result)
	// Output: [1 2 3 4 5]
}

func ExampleSplitAt() {
	slice := []int{1, 2, 3, 4, 5}
	left, right := gslice.SplitAt(slice, 2)
	fmt.Printf("%v, %v\n", left, right)
	// Output: [1 2], [3 4 5]
}
