package main

import "testing"

func Test_maxSubArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "test single value arrays",
			nums: []int{5},
			want: 5,
		},
		{
			name: "test array with all positive ints",
			nums: []int{1, 2, 5, 4, 2, 4},
			want: 18,
		},
		{
			name: "test array with postive and negative ints",
			nums: []int{1, -2, -5, 4, -2, 4},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray(tt.nums); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
