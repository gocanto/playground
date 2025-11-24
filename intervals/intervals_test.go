package intervals

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		want      [][]int
	}{
		{
			name:      "Example from main.go",
			intervals: [][]int{{1, 3}, {2, 4}, {6, 8}, {7, 9}},
			want:      [][]int{{1, 4}, {6, 9}},
		},
		{
			name:      "No overlapping intervals",
			intervals: [][]int{{1, 2}, {3, 4}, {5, 6}},
			want:      [][]int{{1, 2}, {3, 4}, {5, 6}},
		},
		{
			name:      "All overlapping intervals",
			intervals: [][]int{{1, 5}, {2, 6}, {3, 7}},
			want:      [][]int{{1, 7}},
		},
		{
			name:      "Intervals with same start and end",
			intervals: [][]int{{1, 3}, {1, 3}, {2, 4}},
			want:      [][]int{{1, 4}},
		},
		{
			name:      "Single interval",
			intervals: [][]int{{1, 5}},
			want:      [][]int{{1, 5}},
		},
		{
			name:      "Empty input",
			intervals: [][]int{},
			want:      [][]int{},
		},
		{
			name:      "Intervals already merged and sorted",
			intervals: [][]int{{1, 4}, {5, 6}, {7, 8}},
			want:      [][]int{{1, 4}, {5, 6}, {7, 8}},
		},
		{
			name:      "Intervals with negative numbers",
			intervals: [][]int{{-10, -5}, {-6, -1}, {0, 5}},
			want:      [][]int{{-10, -1}, {0, 5}},
		},
		{
			name:      "Intervals with zero",
			intervals: [][]int{{0, 0}, {0, 1}, {1, 2}},
			want:      [][]int{{0, 2}},
		},
		{
			name:      "Complex overlapping",
			intervals: [][]int{{1, 4}, {0, 4}},
			want:      [][]int{{0, 4}},
		},
		{
			name:      "Another complex overlapping",
			intervals: [][]int{{1, 4}, {0, 0}},
			want:      [][]int{{0, 0}, {1, 4}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Merge(tt.intervals)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
