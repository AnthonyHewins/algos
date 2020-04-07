package main

import (
	"fmt"
	"testing"
)

type test struct {
	array []int
}

func (t test) Len() int {
	return len(t.array)
}

func (t test) Less(i, j int) bool {
	return t.array[i] <= t.array[j]
}

func (t test) Swap(i, j int) {
	t.array[i], t.array[j] = t.array[j], t.array[i]
}

func TestToMinHeap(t *testing.T) {
	if err := ToMinHeap(test{ []int{} }, 1); err == nil {
		t.Errorf("should have gotten an error about stop being too large")
	}

	if err := ToMinHeap(test{ []int{} }, -1); err == nil {
		t.Errorf("should have gotten an error for negative stop value")
	}

	if err := ToMinHeap(test{ []int{1} }, 0); err != nil {
		t.Errorf("shouldn't have gotten an error for a no-op")
	}

	scaffoldMin(
		t,
		[]int{5, 8, 7, 26, 1, 80},
		0,
	)
	scaffoldMin(
		t,
		[]int{5, 4, 3, 2, 1},
		0,
	)
	scaffoldMin(
		t,
		[]int{6464,654,6894,6514,684,6841,654,8964,531,35489,456,168,4684,687,464,1351,684,894,651,864,681,354,864,684,68468,465,16,468,16,132,035,168,168,1651,6810,35,68,06,68,126,06,82168,165,0,654086,06,512,98612,6512,864,641,68,681,651,684,68,486,465,3135,13,5202,0,3520,50,505,056,06,1268,568,468,46,816,53,68,62,68,168,168,16},
		0,
	)
	scaffoldMin(
		t,
		[]int{45,4,8,1,51,54,81,51,51,51,84,84,6,23,2,65,94,84,562,65,9,94,51,5,48,7,92,651,584,84,89,16,218,48,484,84,84,84,84,848,48,48,11111,0,1,21,8,0},
		0,
	)
	scaffoldMin(
		t,
		[]int{45,4,8,1,51,54,81,51,51,51,84,84,6,23,2,65,94,84,562,65,9,94,51,5,48,7,92,651,584,84,89,16,218,48,484,84,84,84,84,848,48,48,11111,0,1,21,8,0},
		8,
	)

	size := 10
	for shift := 0; shift < size; shift++ {
		for stop := 2; stop <= size; stop++ {
			arr := make([]int, size, size)

			for i := range arr { arr[i] = (i + shift) % size }

			scaffoldMin(t, arr, stop)
		}
	}
}

func TestToMaxHeap(t *testing.T) {
	if err := ToMaxHeap(test{ []int{} }, 1); err == nil {
		t.Errorf("should have gotten an error about stop being too large")
	}

	if err := ToMaxHeap(test{ []int{} }, -1); err == nil {
		t.Errorf("should have gotten an error for negative stop value")
	}

	if err := ToMaxHeap(test{ []int{1} }, 0); err != nil {
		t.Errorf("shouldn't have gotten an error for a no-op")
	}

	scaffoldMax(
		t,
		[]int{5, 8, 7, 26, 1, 80},
		0,
	)
	scaffoldMax(
		t,
		[]int{5, 4, 3, 2, 1},
		0,
	)
	scaffoldMax(
		t,
		[]int{6464,654,6894,6514,684,6841,654,8964,531,35489,456,168,4684,687,464,1351,684,894,651,864,681,354,864,684,68468,465,16,468,16,132,035,168,168,1651,6810,35,68,06,68,126,06,82168,165,0,654086,06,512,98612,6512,864,641,68,681,651,684,68,486,465,3135,13,5202,0,3520,50,505,056,06,1268,568,468,46,816,53,68,62,68,168,168,16},
		0,
	)
	scaffoldMax(
		t,
		[]int{45,4,8,1,51,54,81,51,51,51,84,84,6,23,2,65,94,84,562,65,9,94,51,5,48,7,92,651,584,84,89,16,218,48,484,84,84,84,84,848,48,48,11111,0,1,21,8,0},
		0,
	)
	scaffoldMax(
		t,
		[]int{45,4,8,1,51,54,81,51,51,51,84,84,6,23,2,65,94,84,562,65,9,94,51,5,48,7,92,651,584,84,89,16,218,48,484,84,84,84,84,848,48,48,11111,0,1,21,8,0},
		8,
	)

	size := 10
	for shift := 0; shift < size; shift++ {
		for stop := 2; stop <= size; stop++ {
			arr := make([]int, size, size)

			for i := range arr { arr[i] = (i + shift) % size }

			scaffoldMax(t, arr, stop)
		}
	}
}

func scaffoldMin(t *testing.T, inp []int, stop int) {
	q1 := test { inp }
	if err := ToMinHeap(q1, stop); err != nil {
		t.Errorf("error with ToMinHeap: %v", err)
	}

	for i := stop - 1; i >= 1; i-- {
		parent := (i - 1) / 2
		if q1.array[parent] > q1.array[i] {
			t.Errorf(
				fmt.Sprintf(
					"heap violation: parent %d > child %d for array\n%v",
					parent,
					i,
					q1.array[:stop],
				),
			)
		}
	}
}

func scaffoldMax(t *testing.T, inp []int, stop int) {
	q1 := test { inp }
	if err := ToMaxHeap(q1, stop); err != nil {
		t.Errorf("error with ToMaxHeap: %v", err)
	}

	for i := stop - 1; i >= 1; i-- {
		parent := (i - 1) / 2
		if q1.array[parent] < q1.array[i] {
			t.Errorf(
				fmt.Sprintf(
					"heap violation: parent %d > child %d for array\n%v",
					parent,
					i,
					q1.array[:stop],
				),
			)
		}
	}
}
