package main

import (
	"sort"
)

// I believe this is n^2*log(n), there is an log(n) solution
// but sort.Slice() is only in Go 1.8 and I didn't want to implement merge sort
// on arrays
func merge(intervals [][]int) [][]int {
	// go is manual on generic sorts so let's start with a hash
	m := make(map[int]int)
	var keys []int
	var output [][]int

	for _, interval := range intervals {
		// [0] = start, [1] = end
		if _, ok := m[interval[0]]; ok {
			// interesting thing here is if we have start vals that are the same
			// then we can combine intervals in our loop
			if m[interval[0]] < interval[1] {
				// made a dumb mistake here where I was calling the map for interval[1] instead of the value
				m[interval[0]] = interval[1]
			}
			continue
		}

		m[interval[0]] = interval[1]
		keys = append(keys, interval[0])
	}

	sort.Ints(keys)
	var currentPair []int

	for i, key := range keys {
		if i == 0 {
			currentPair = []int{key, m[key]}
		}

		if key <= currentPair[1] {
			// missed an edge case here where 1 range can be a subset of another
			if currentPair[1] < m[key] {
				currentPair = []int{currentPair[0], m[key]}
			}
		} else {
			output = append(output, currentPair)
			currentPair = []int{key, m[key]}
		}

		if i == len(keys)-1 {
			output = append(output, currentPair)
		}
	}

	return output
}
