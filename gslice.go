package gslice

import (
	"errors"
)

var (
	ErrOutOfRange = errors.New("index out of range")
)

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

