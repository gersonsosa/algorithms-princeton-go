package algorithms

import "testing"

func TestFind(t *testing.T) {
	type args struct {
		x int
		a []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"should find in 3lgn",
			args{6, []int{1,3,5,7,9,11,20,10,8,6,4,-1}},
			9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Find(tt.args.x, tt.args.a); got != tt.want {
				t.Errorf("Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFind2lgn(t *testing.T) {
	type args struct {
		x int
		a []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"should find in 3lgn",
			args{6, []int{1,3,5,9,15,12,10,8,6,4,2,-1}},
			9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Find2lgn(tt.args.x, tt.args.a); got != tt.want {
				t.Errorf("Find() = %v, want %v", got, tt.want)
			}
		})
	}
}
