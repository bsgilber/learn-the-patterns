package main

import "testing"

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "length 2 with duplicate",
			args: args{
				nums: []int{1, 1},
			},
			want: 1,
		},
		{
			name: "single duplicate",
			args: args{
				nums: []int{1, 1, 2},
			},
			want: 2,
		},
		{
			name: "multiple duplicates",
			args: args{
				nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.nums); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeDuplicatesBetter(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "length 2 with duplicate",
			args: args{
				nums: []int{1, 1},
			},
			want: 1,
		},
		{
			name: "single duplicate",
			args: args{
				nums: []int{1, 1, 2},
			},
			want: 2,
		},
		{
			name: "multiple duplicates",
			args: args{
				nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicatesBetter(tt.args.nums); got != tt.want {
				t.Errorf("removeDuplicatesBetter() = %v, want %v", got, tt.want)
			}
		})
	}
}
