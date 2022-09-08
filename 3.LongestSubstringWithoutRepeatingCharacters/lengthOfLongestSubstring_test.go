package main

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Ex1",
			args: args{s: "abcabcbb"},
			want: 3,
		},
		{
			name: "Ex2",
			args: args{s: " "},
			want: 1,
		},
		{
			name: "Ex3",
			args: args{s: "pwwkew"},
			want: 3,
		},
		{
			name: "Ex4",
			args: args{s: "au"},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
