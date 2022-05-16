package main

import "testing"

func Test_threeSumSmaller(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "full example with 2 solutions",
			args: args{
				nums:   []int{-2, 0, 1, 3},
				target: 2,
			},
			want: 2,
		},
		{
			name: "empty input",
			args: args{
				nums:   []int{},
				target: 0,
			},
			want: 0,
		},
		{
			name: "single input that is the same as target",
			args: args{
				nums:   []int{0},
				target: 0,
			},
			want: 0,
		},
		{
			name: "multiple index should still zero out because i<j<k",
			args: args{
				nums:   []int{0, 0, 0, 0, 0},
				target: 0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSumSmaller(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("threeSumSmaller() = %v, want %v", got, tt.want)
			}
		})
	}
}
