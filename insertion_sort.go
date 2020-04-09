package main

import (
	"sort"
)

func InsertionSort(data sort.Interface) {
	is(data, 0, data.Len() - 1)
}

func is(data sort.Interface, start, end int) {
	for i := start + 1; i <= end; i++ {
		ptr := i
		for j := ptr - 1; j >= 0; j, ptr = j - 1, ptr - 1 {
			if data.Less(j, ptr) { break }
			data.Swap(j, ptr)
		}
	}
}
