package c5

import (
	"sort"
	"testing"
)

func TestSlice(t *testing.T) {
	values := []int32{42, 9, 101, 95, 27, 25}

	Slice(values, func(i, j int) bool {
		return values[i] < values[j]
	})

	isSorted := sort.SliceIsSorted(values, func(i, j int) bool {
		return values[i] < values[j]
	})
	if !isSorted {
		t.Fatal("should be sorted")
	}
}
