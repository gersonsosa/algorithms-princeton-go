package stackmax

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *StackMax
	}{
		{"should create a new StackMax struct", &StackMax{make([]int, 0, 10), -1, -1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStackMax_findMax(t *testing.T) {
	type fields struct {
		Items []int
		max   int
		last  int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"should find max(5)", fields{[]int{3, 2, 0, 5, 1}, 3, 4}, 5},
		{"should find max(11)", fields{[]int{10, 9, 8, 7, 11}, 4, 4}, 11},
		{"should find max(0)", fields{[]int{0, 0, 0, 0, 0}, 4, 4}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stackMax := StackMax{
				Items: tt.fields.Items,
				max:   tt.fields.max,
				last:  tt.fields.last,
			}
			if got := stackMax.findMax(); got != tt.want {
				t.Errorf("StackMax.findMax() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStackMax_Push(t *testing.T) {
	type fields struct {
		Items []int
		max   int
		last  int
	}
	type args struct {
		i int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"should push element(5) to the top of the stack", fields{[]int{3, 2, 0, 5, 1}, 3, 4}, args{0}},
		{"should push element(11) to the top of the stack", fields{[]int{10, 9, 8, 7, 11}, 4, 4}, args{1}},
		{"should push element(0) to the top of the stack", fields{[]int{0, 0, 0, 0, 0}, 4, 4}, args{123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stackMax := StackMax{
				Items: tt.fields.Items,
				max:   tt.fields.max,
				last:  tt.fields.last,
			}
			stackMax.Push(tt.args.i)
			if got := stackMax.Pop(); got != tt.args.i {
				t.Errorf("StackMax.Pop() = %v, want %v", got, tt.args.i)
			}
		})
	}
}

func TestStackMax_Pop(t *testing.T) {
	type fields struct {
		Items []int
		max   int
		last  int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"should pop last element(1)", fields{[]int{3, 2, 0, 5, 1}, 3, 4}, 1},
		{"should pop last element(11)", fields{[]int{10, 9, 8, 7, 11}, 4, 4}, 11},
		{"should pop last element(0)", fields{[]int{0, 0, 0, 0, 0}, 4, 4}, 0},
		{"should pop last element(7)", fields{[]int{7}, 0, 0}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stackMax := StackMax{
				Items: tt.fields.Items,
				max:   tt.fields.max,
				last:  tt.fields.last,
			}
			if got := stackMax.Pop(); got != tt.want {
				t.Errorf("StackMax.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStackMax_shrinkIfNecessary(t *testing.T) {
	type fields struct {
		Items []int
		max   int
		last  int
	}
	tests := []struct {
		name   string
		fields fields
		want int
	}{
		{"should shrink array when is 3/4 empty and pop is invoked", fields{make([]int, 3, 8), 0, 2}, 4},
		{"should shrink array when is 3/4 empty and pop is invoked", fields{make([]int, 6, 20), 0, 5}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stackMax := StackMax{
				Items: tt.fields.Items,
				max:   tt.fields.max,
				last:  tt.fields.last,
			}
			stackMax.Pop()
			if cap(stackMax.Items) > len(stackMax.Items) * 2 {
				t.Errorf("StackMax.shrinkIfNecessary() cap = %v, want %v", cap(stackMax.Items), tt.want)
			}
		})
	}
}

func TestStackMax_searchNewMax(t *testing.T) {
	type fields struct {
		Items []int
		max   int
		last  int
	}
	tests := []struct {
		name   string
		fields fields
		want int
	}{
		{"should return max value from array", fields{[]int{3, 2, 0, 5, 10}, 4, 4}, 5},
		{"should return max value from array", fields{[]int{20, 15, 10, 2, 28}, 4, 4}, 20},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stackMax := StackMax{
				Items: tt.fields.Items,
				max:   tt.fields.max,
				last:  tt.fields.last,
			}
			stackMax.Pop()
			if stackMax.Items[stackMax.max] != tt.want {
				t.Errorf("StackMax.searchNewMax() max = %v, want %v", stackMax.Items[stackMax.max], tt.want)
			}
		})
	}
}
