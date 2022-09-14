package main

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ex1",
			args: args{s: "12345"},
			want: 12345,
		},
		{
			name: "ex2",
			args: args{s: "4193 with words"},
			want: 4193,
		},
		{
			name: "ex3",
			args: args{s: "20000000000000000000"},
			want: 2147483647,
		},
		{
			name: "ex4",
			args: args{s: "words and 987"},
			want: 0,
		},
		{
			name: "ex5",
			args: args{s: "-91283472332"},
			want: -2147483648,
		},
		{
			name: "ex6",
			args: args{s: "  0000000000012345678"},
			want: 12345678,
		},
		{
			name: "ex7",
			args: args{s: "00000-42a1234"},
			want: 0,
		},
		{
			name: "ex8",
			args: args{s: "-000000000000001"},
			want: -1,
		},
		{
			name: "ex9",
			args: args{s: "  -0012a42"},
			want: -12,
		},
		{
			name: "ex10",
			args: args{s: "    0000000000000   "},
			want: 0,
		},
		{
			name: "ex11",
			args: args{s: "-5-"},
			want: -5,
		},
		{
			name: "ex12",
			args: args{s: "-13+8"},
			want: -13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.s); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
