package main

import (
	"reflect"
	"testing"
)

func Test_threeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "multiple sums to 0",
			args: args{
				nums: []int{-1, 0, 1, 2, -1, 4},
			},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name: "empty input list",
			args: args{
				nums: []int{},
			},
			want: [][]int{},
		},
		{
			name: "less than 3 elements",
			args: args{
				nums: []int{3},
			},
			want: [][]int{},
		},
		{
			name: "no solution",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
			},
			want: [][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
