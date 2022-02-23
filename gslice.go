package gslice

import (
	"errors"
)

var (
	ErrOutOfRange = errors.New("index out of range")
)

// min returns the smallest of the given values.
func min(a int, bs ...int) int {
	smallest := a
	for _, b := range bs {
		if b < smallest {
			smallest = b
		}
	}
	return smallest
}

// max returns the largest of the given values.
func max(a int, bs ...int) int {
	largest := a
	for _, b := range bs {
		if b > largest {
			largest = b
		}
	}
	return largest
}

// clipIndex returns the given index, clipped to the range [minI, maxI].
// It can be used to ensure that an index is within the bounds of a slice.
func clipIndex(i, minI, maxI int) int {
	return min(max(i, minI), maxI)
}

// AppendAt inserts the element (or elements) at the given index.
// If the index is out of range, an error is returned.
//
// AppendAt is similar to the built-in append function, except that
// it will insert the element(s) at the specified index.
//
// TODO: Add negative index support
func AppendAt[T any](slice []T, index int, elems ...T) ([]T, error) {
	if index < 0 || index > len(slice) {
		return nil, ErrOutOfRange
	}
	return append(
		slice[:index],
		append(
			elems,
			slice[index:]...,
		)...,
	), nil
}

// MustAppendAt inserts the given value at the given index.
// It wraps AppendAt and panics if an error is returned – for
// example, if the index is out of range.
// 
// For more information, see the documentation for AppendAt.
func MustAppendAt[T any](slice []T, index int, elems ...T) []T {
	result, err := AppendAt(slice, index, elems...)
	if err != nil {
		panic(err)
	}
	return result
}

// SplitAt splits the given slice at the given index.
// If the index is out of range (to the left or right), no error
// will be returned, instead that split will be empty.
func SplitAt[T any](slice []T, index int) ([]T, []T) {
	s := clipIndex(index, 0, len(slice))
	return slice[:s], slice[s:]
}

