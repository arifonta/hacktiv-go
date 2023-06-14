package main

import "testing"

func Test_fib(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Should be positive number",
			args: args{
				num: -1,
			},
			want: "0",
		},
		// TODO: Add test cases
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fib(tt.args.num)
		})
	}
}
