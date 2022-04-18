package main

import "testing"

func Test_minSubArrayLen(t *testing.T) {
	type args struct {
		target int
		nums   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test array with length 1 that does not contain target",
			args: args{
				target: 4,
				nums:   []int{3},
			},
			want: 0,
		},
		{
			name: "test array with length 1 that does contain target",
			args: args{
				target: 4,
				nums:   []int{4},
			},
			want: 1,
		},
		{
			name: "test array that does not contain subArray = target",
			args: args{
				target: 18,
				nums:   []int{1, 2, 3, 4, 5},
			},
			want: 0,
		},
		{
			name: "test array that contains 1 subArray = target",
			args: args{
				target: 7,
				nums:   []int{1, 2, 3, 4, 5},
			},
			want: 2,
		},
		{
			name: "test array that contains >1 subArrays = target",
			args: args{
				target: 5,
				nums:   []int{1, 2, 3, 4, 5},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubArrayLen(tt.args.target, tt.args.nums); got != tt.want {
				t.Errorf("minSubArrayLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
