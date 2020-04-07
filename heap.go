package main

import (
	"fmt"
	"sort"
)

// Turn argument into a heap in-place
// Give a nonzero value to tell the algorithm to stop at value "stop" (exclusively)
// in case you don't want the entire thing to be turned into a heap
func ToMaxHeap(data sort.Interface, stop int) error {
	return heapify(data, stop, true)
}

func ToMinHeap(data sort.Interface, stop int) error {
	return heapify(data, stop, false)
}

func heapify(data sort.Interface, stop int, isMaxHeap bool) error {
	n := data.Len()

	if stop > n {
		return fmt.Errorf(
			"ToMinHeap got a stop too large for array of length %d: %d", stop, n,
		)
	} else if stop < 0 {
		return fmt.Errorf("got negative index: %d", stop)
	}

	// No work to be done
	if n <= 1 || stop == 1 { return nil }

	// Special case which we populate as n as a convenience to the user
	if stop == 0 { stop = n	}

	// 1 < stop < n
	for i := 1; i < stop; i++ {
		// ptr to keep track of where this element is as it bubbles up
		ptr := i

		for parent := (ptr - 1) / 2; (data.Less(ptr, parent) != isMaxHeap) && ptr > 0; parent = (ptr - 1) / 2 {
			data.Swap(ptr, parent)
			ptr = parent
		}
	}

	return nil
}
