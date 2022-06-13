package main

import "testing"

func Test_testNb(t *testing.T) {
	type args struct {
		nb  float64
		max int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "5.0",
			args: args{nb: 5.0, max: 10},
			want: 1,
		},
		{
			name: "6.0",
			args: args{nb: 6.0, max: 10},
			want: 0,
		},
		{
			name: "6263.0",
			args: args{nb: 6263.0, max: 7000},
			want: 1,
		},
		{
			name: "3212.0",
			args: args{nb: 3212.0, max: 4000},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := testNb(tt.args.nb, tt.args.max); got != tt.want {
				t.Errorf("testNb() = %v, want %v", got, tt.want)
			}
		})
	}
}
