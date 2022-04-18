package main

import "testing"

func Test_lengthOfLongestSubstringKDistinct(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test string of length 1 with k=1",
			args: args{
				s: "a",
				k: 1,
			},
			want: 1,
		},
		{
			name: "test string with k=2, only 1 char appears twice",
			args: args{
				s: "ebeap",
				k: 2,
			},
			want: 3,
		},
		{
			name: "test string with k=2, multiple chars repeat",
			args: args{
				s: "ebeapp",
				k: 2,
			},
			want: 3,
		},
		{
			name: "test string with k=3, multiple chars repeat",
			args: args{
				s: "ebeappep",
				k: 3,
			},
			want: 6,
		},
		{
			name: "test string with k=3, multiple chars repeat",
			args: args{
				s: "ebapppe",
				k: 3,
			},
			want: 5,
		},
		{
			name: "test string with k=3, multiple chars repeat",
			args: args{
				s: "ccaabbb",
				k: 2,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstringKDistinct(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("lengthOfLongestSubstringKDistinct() = %v, want %v", got, tt.want)
			}
		})
	}
}
