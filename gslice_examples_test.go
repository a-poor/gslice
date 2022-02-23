package gslice_test

import (
	"log"
	"reflect"

	"github.com/a-poor/gslice"
)

func ExampleAppendAt() {
	start := []int{1, 2, 4, 5}
	expected := []int{1, 2, 3, 4, 5}

	// Add a single value to the middle of the slice
	result, err := gslice.AppendAt(start, 2, 3)

	if err != nil {
		log.Panic(err)
	}
	if !reflect.DeepEqual(result, expected) {
		log.Panicf("AppendAt(%v, 2, 3) = %v, want %v", start, result, expected)
	}
}

func ExampleAppendAt_multiple() {
	start := []int{1, 5}
	expected := []int{1, 2, 3, 4, 5}

	// Add multiple elements at once
	result, err := gslice.AppendAt(start, 2, 2, 3, 4)

	if err != nil {
		log.Panic(err)
	}
	if !reflect.DeepEqual(result, expected) {
		log.Panicf("AppendAt(%v, 2, 3) = %v, want %v", start, result, expected)
	}
}
