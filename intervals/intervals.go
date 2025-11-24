package intervals

import (
	"sort"
)

// go run main.go
//    Input: [[1 3] [2 4] [6 8] [7 9]]
//    Output: [[1 4] [6 9]]

// Merge takes a slice of intervals (each represented as a slice of 2 ints)
// and returns a new slice of non-overlapping merged intervals.
// Time Complexity: O(N log N) due to sorting.
// Space Complexity: O(N) for the output slice.
func Merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	// 1. Sort intervals by start time.
	//    We use sort.Slice for convenience. ideally in a hot path,
	//    implementing sort.Interface on a custom type might be slightly faster
	//    due to less reflection overhead, but Slice is highly optimised in modern Go.
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var result [][]int

	// Initialise with the first interval
	// Note: We copy the interval to avoid mutating the original underlying array
	//       if we modify result[len-1][1] later.
	current := []int{intervals[0][0], intervals[0][1]}
	result = append(result, current)

	for _, interval := range intervals[1:] {
		lastMerged := result[len(result)-1]

		// 2. Compare the current interval with the last merged interval
		if interval[0] <= lastMerged[1] {
			// Overlap detected: Merge them.
			// Update the end time of the last merged interval.
			if interval[1] > lastMerged[1] {
				lastMerged[1] = interval[1]
			}
		} else {
			// No overlap: Add the current interval to the result.
			// We create a new slice to ensure safety.
			newInterval := []int{interval[0], interval[1]}
			result = append(result, newInterval)
		}
	}

	return result
}
