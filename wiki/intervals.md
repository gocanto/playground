# Interval Merging in Go

This project provides a Go implementation for merging overlapping intervals. It demonstrates an efficient algorithm to consolidate a list of intervals into a set of non-overlapping intervals.

## Problem Description

Given a collection of intervals, merge all overlapping intervals.

### Example

**Input:** `[[1, 3], [2, 4], [6, 8], [7, 9]]`
**Output:** `[[1, 4], [6, 9]]`

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

You need to have Go installed on your system (version 1.25.3 or higher).

### Running the Example

To see the interval merging in action with a predefined test case:

```bash
go run main.go
```

Expected Output:

```
Input:  [[1 3] [2 4] [6 8] [7 9]]
Output: [[1 4] [6 9]]
```

### Running Tests

To execute the unit tests for the interval merging logic:

```bash
make tests
```

### Formatting Code

To format the Go source code:

```bash
make format
```

## Algorithm Details

The `Merge` function in `intervals/intervals.go` implements the following approach:

1.  **Sort Intervals:** The intervals are first sorted based on their start times.
2.  **Iterate and Merge:** The algorithm then iterates through the sorted intervals, merging an interval with the previous one if they overlap, or adding it as a new non-overlapping interval if no overlap exists.

### Complexity

*   **Time Complexity:** O(N log N) due to the initial sorting of intervals.
*   **Space Complexity:** O(N) for storing the merged intervals in the worst case (e.g., no overlaps, so all intervals are added to the result).

## Built With

*   [Go—](https://golang.org/)The programming language used.