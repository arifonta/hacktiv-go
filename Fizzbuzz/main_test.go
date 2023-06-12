package main

import "testing"

func Test_fizzBuzz(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "should get fizz",
			args: args{
				i: 3,
			},
			want: "fizz",
		},
		{
			name: "should get buzz",
			args: args{
				i: 5,
			},
			want: "buzz",
		},
		{
			name: "should get fizzbuzz",
			args: args{
				i: 15,
			},
			want: "fizzbuzz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fizzBuzz(tt.args.i); got != tt.want {
				t.Errorf("fizzBuzz() = %v, want %v", got, tt.want)
			}
		})
	}
}
