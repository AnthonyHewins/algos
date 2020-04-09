package main

import (
	"sort"
	"testing"
	"math/rand"
)

const (
	IS_SIZE = 300
	QS_SIZE = 2000
)


func TestQuicksort(t *testing.T) {
	sortTest(t, QS_SIZE, Quicksort)
}

func TestInsertionSort(t *testing.T) {
	sortTest(t, IS_SIZE, InsertionSort)
}

func sortTest(t *testing.T, size int, algo func(sort.Interface)) {
	for currentSize := 0; currentSize < size; currentSize++ {
		input := make([]int, currentSize, currentSize)

		// all shifted permutations of Z_n
		for shift := 0; shift < currentSize; shift++ {
			for i := 0; i < currentSize; i++ {
				input[i] = i + shift
			}

			driver(t, input, Quicksort)
		}

		// random arrays
		for randCount := 0; randCount < 10; randCount++ {
			for i := 0; i < currentSize; i++ {
				input[i] = rand.Intn(1000)
			}

			driver(t, input, Quicksort)
		}
	}
}

func driver(t *testing.T, arr []int, algo func(sort.Interface)) {
	x := test{arr}
	algo(x)

	if !sort.IsSorted(x) {
		t.Errorf("array was not sorted: %v", x.array)
	}
}
