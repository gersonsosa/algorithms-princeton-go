package main

import "testing"

func TestPermutation(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"should return true for permutable arrays",
			args{
				[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				[]int{9, 8, 3, 6, 7, 5, 2, 4, 1, 0},
			},
			true,
		}, {
			"should return false for different size arrays",
			args{
				[]int{0},
				[]int{9, 8, 3, 6, 7, 5, 2, 4, 1, 0},
			},
			false,
		}, {
			"should return false for different arrays",
			args{
				[]int{50, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				[]int{9, 8, 3, 6, 7, 5, 2, 4, 1, 0},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPermutation(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("IsPermutation(%v, %v) = %v, want %v", tt.args.a, tt.args.b, got, tt.want)
			}
		})
	}
}
