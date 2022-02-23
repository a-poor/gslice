package gslice_test

import (
	"reflect"
	"testing"

	"github.com/a-poor/gslice"
)

func TestAppendAt(t *testing.T) {
	t.Run("insert-into-middle", func(t *testing.T) {
		start := []int{1, 2, 4, 5}
		index := 2
		add := 3
		expected := []int{1, 2, 3, 4, 5}

		result, err := gslice.AppendAt(start, index, add)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf(
				"AppendAt(%v, %v, %v) = %v, want %v",
				start, index, add, result, expected,
			)
		}
	})

	t.Run("insert-at-start", func(t *testing.T) {
		start := []int{1, 2, 4, 5}
		index := 0
		add := 3
		expected := []int{3, 1, 2, 4, 5}

		result, err := gslice.AppendAt(start, index, add)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf(
				"AppendAt(%v, %v, %v) = %v, want %v",
				start, index, add, result, expected,
			)
		}
	})

	t.Run("insert-at-end", func(t *testing.T) {
		start := []int{1, 2, 4, 5}
		index := len(start)
		add := 3
		expected := []int{1, 2, 4, 5, 3}

		result, err := gslice.AppendAt(start, index, add)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf(
				"AppendAt(%v, %v, %v) = %v, want %v",
				start, index, add, result, expected,
			)
		}
	})

	t.Run("shouldnt-insert-before-beginning", func(t *testing.T) {
		start := []int{1, 2, 4, 5}
		index := -1
		add := 3
		expectedErr := gslice.ErrOutOfRange

		result, err := gslice.AppendAt(start, index, add)
		if err == nil {
			t.Errorf(
				"AppendAt(%v, %v, %v) = %v, want error %v",
				start, index, add, result, expectedErr,
			)
		}
		if err != expectedErr {
			t.Errorf("Got error: %q, Expected error: %q", err, expectedErr)
		}
	})

	t.Run("shouldnt-insert-after-end", func(t *testing.T) {
		start := []int{1, 2, 4, 5}
		index := 100
		add := 3
		expectedErr := gslice.ErrOutOfRange

		result, err := gslice.AppendAt(start, index, add)
		if err == nil {
			t.Errorf(
				"AppendAt(%v, %v, %v) = %v, want error %v",
				start, index, add, result, expectedErr,
			)
		}
		if err != expectedErr {
			t.Errorf("Got error: %q, Expected error: %q", err, expectedErr)
		}
	})
}

func TestAppendAtWithMultipleValues(t *testing.T) {
	t.Run("insert-into-middle", func(t *testing.T) {
		start := []int{1, 2, 5}
		index := 2
		add := []int{3, 4}
		expected := []int{1, 2, 3, 4, 5}

		result, err := gslice.AppendAt(start, index, add...)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf(
				"AppendAt(%v, %v, %v) = %v, want %v",
				start, index, add, result, expected,
			)
		}
	})

	t.Run("insert-at-start", func(t *testing.T) {
		start := []int{3, 4, 5}
		index := 0
		add := []int{1, 2}
		expected := []int{1, 2, 3, 4, 5}

		result, err := gslice.AppendAt(start, index, add...)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf(
				"AppendAt(%v, %v, %v) = %v, want %v",
				start, index, add, result, expected,
			)
		}
	})

	t.Run("insert-at-end", func(t *testing.T) {
		start := []int{1, 2, 3, 4}
		index := len(start)
		add := []int{5}
		expected := []int{1, 2, 3, 4, 5}

		result, err := gslice.AppendAt(start, index, add...)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf(
				"AppendAt(%v, %v, %v) = %v, want %v",
				start, index, add, result, expected,
			)
		}
	})

	t.Run("shouldnt-insert-before-beginning", func(t *testing.T) {
		start := []int{1, 2, 4, 5}
		index := -1
		add := []int{1, 2}
		expectedErr := gslice.ErrOutOfRange

		result, err := gslice.AppendAt(start, index, add...)
		if err == nil {
			t.Errorf(
				"AppendAt(%v, %v, %v) = %v, want error %v",
				start, index, add, result, expectedErr,
			)
		}
		if err != expectedErr {
			t.Errorf("Got error: %q, Expected error: %q", err, expectedErr)
		}
	})

	t.Run("shouldnt-insert-after-end", func(t *testing.T) {
		start := []int{1, 2, 4, 5}
		index := 100
		add := []int{1, 2}
		expectedErr := gslice.ErrOutOfRange

		result, err := gslice.AppendAt(start, index, add...)
		if err == nil {
			t.Errorf(
				"AppendAt(%v, %v, %v) = %v, want error %v",
				start, index, add, result, expectedErr,
			)
		}
		if err != expectedErr {
			t.Errorf("Got error: %q, Expected error: %q", err, expectedErr)
		}
	})
}

func TestSplitAt(t *testing.T) {
	testCases := []struct {
		name  string
		slice []int
		index int
		left  []int
		right []int
	}{
		{
			name:  "split-at-middle-even",
			slice: []int{1, 2, 4, 5},
			index: 2,
			left:  []int{1, 2},
			right: []int{4, 5},
		},
		{
			name:  "split-at-middle-uneven",
			slice: []int{1, 2, 3, 4, 5},
			index: 3,
			left:  []int{1, 2, 3},
			right: []int{4, 5},
		},
		{
			name:  "split-at-front",
			slice: []int{1, 2, 3, 4, 5},
			index: 0,
			left:  []int{},
			right: []int{1, 2, 3, 4, 5},
		},
		{
			name:  "split-at-end",
			slice: []int{1, 2, 3, 4, 5},
			index: 5,
			left:  []int{1, 2, 3, 4, 5},
			right: []int{},
		},
		{
			name:  "split-before-beginning",
			slice: []int{1, 2, 3, 4, 5},
			index: -10,
			left:  []int{},
			right: []int{1, 2, 3, 4, 5},
		},
		{
			name:  "split-after-end",
			slice: []int{1, 2, 3, 4, 5},
			index: 10,
			left:  []int{1, 2, 3, 4, 5},
			right: []int{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Perform the split
			left, right := gslice.SplitAt(tc.slice, tc.index)

			// Check the results
			leftMatch := reflect.DeepEqual(left, tc.left)
			rightMatch := reflect.DeepEqual(right, tc.right)

			if !leftMatch || !rightMatch {
				t.Errorf(
					"SplitAt(%v, %v) = %v, %v, want %v, %v",
					tc.slice, tc.index, left, right, tc.left, tc.right,
				)
			}
		})
	}
}
