package main

import "testing"

/*
Example 1:

Input: x = 123
Output: 321
Example 2:

Input: x = -123
Output: -321
Example 3:

Input: x = 120
Output: 21

*/

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ex1",
			args: args{x: 123},
			want: 321,
		},
		{
			name: "ex2",
			args: args{x: 120},
			want: 21,
		},
		{
			name: "ex3",
			args: args{x: -123},
			want: -321,
		},
		{
			name: "ex4",
			args: args{x: 1534236469},
			want: 0,
		},
		{
			name: "ex5",
			args: args{x: -2147483648},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
