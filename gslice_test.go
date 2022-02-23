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
