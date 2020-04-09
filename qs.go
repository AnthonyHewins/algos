package main

import (
	"sort"
	"sync"
)

// These constants need to make sense; check the qs logic to make sure
// that after a change they won't break assumptions that are super
// helpful in making it fast
const (
	// Experimental constant for insertion sort.
	// At certain array lengths it will be more optimal to just
	// use an O(n^2) sort because it performs better on average
	// than using nlog(n) sorts; should be a small value
	insertionSortThreshold = 7

	// These constants are used AFTER the insertionSortThreshold,
	// meaning these are only used after it's confirmed the length of
	// the slice being sorted is at least (insertionSortThreshold + 1)
	medianOf3Threshold        = 50
	medianOf3TriplesThreshold = 500
)

func Quicksort(data sort.Interface) {
	var qs func(int, int)
	var wg sync.WaitGroup

	qs = func (start, end int) {
		n := end - start
		if n <= 0 {
			wg.Done()
			return
		}

		if n == 1 {
			if data.Less(end, start) { data.Swap(end, start) }
			wg.Done()
			return
		}

		if n <= insertionSortThreshold {
			is(data, start, end)
			wg.Done()
			return
		}

		pivot := pseudoMedian(data, start, end, n)
		data.Swap(pivot, start)

		border := start
		for ptr := start + 1; ptr <= end; ptr++ {
			if data.Less(ptr, start) {
				border++
				data.Swap(ptr, border)
			}
		}

		data.Swap(start, border)

		wg.Add(1)
		go qs(start, border - 1)
		qs(border + 1, end)
	}

	wg.Add(1)
	qs(0, data.Len() - 1)
	wg.Wait()
}

func pseudoMedian(data sort.Interface, start, end, n int) int {
	reduceMedian := func(begin, mid, stop int) int {
		if data.Less(begin, mid) {
			if data.Less(mid,   stop) { return mid }
			if data.Less(begin, stop) { return begin }
			return stop
		} else {
			if data.Less(begin, stop) { return begin }
			if data.Less(mid,   stop) { return mid }
			return stop
		}
	}

	// Median of
	//  V                   V                V
	// [start start+1 ... middle ... last-1 last]
	midpoint := (start + end) / 2
	if n <= medianOf3Threshold {
		return reduceMedian(start, midpoint, end)
	}


	startMedian := reduceMedian(start,        start + 1, start + 2)
	midMedian   := reduceMedian(midpoint - 1, midpoint,  midpoint + 1)
	endMedian   := reduceMedian(end - 2,      end - 1,   end)

	medianOfThree := reduceMedian(startMedian, midMedian, endMedian)

	// Median of
	//  V        V      V           V        V     V            V       V    V
	// [start start+1 start+2... middle-1 middle middle+1 ... last-2 last-1 last]
	if n <= medianOf3TriplesThreshold { return medianOfThree }

	// "quantiles"; in reality they aren't but the var gets the message across
	quantile1 := (start + midpoint) / 2
	quantile3 := (end   + midpoint) / 2

	// You get the idea
	return reduceMedian(
		reduceMedian(quantile1 - 1, quantile1, quantile1 + 1),
		medianOfThree,
		reduceMedian(quantile3 - 1, quantile3, quantile3 + 1),
	)
}
