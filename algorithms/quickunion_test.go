package algorithms

import (
	"reflect"
	"testing"
)

func TestNewQuickUnion(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args   args
		want QuickUnion
	}{
		{"should create quickunion struct", args{10},
		&quickUnion{[]int{0,1,2,3,4,5,6,7,8,9}, []int{1,1,1,1,1,1,1,1,1,1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewQuickUnion(tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewQuickUnion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_quickUnion_root(t *testing.T) {
	type fields struct {
		Items []int
	}
	type args struct {
		i int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{"should find the root", fields{[]int{0,0,0,0,2,2,2,5,6,7}}, args{9}, 0},
		{"should find the root in a deep tree", fields{[]int{0,0,1,2,3,4,5,6,7,8}}, args{9}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := quickUnion{
				Items: tt.fields.Items,
			}
			if got := q.root(tt.args.i); got != tt.want {
				t.Errorf("quickUnion.root() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_quickUnion_Union(t *testing.T) {
	type fields struct {
		Items []int
		subtreeSize []int
	}
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"should join two elements", fields{[]int{0,0,0,0,2,2,2,5,6,9}, []int{9,9,9,9,9,9,9,9,9,0}}, args{0, 9}},
		{"should join two elements", fields{[]int{0,1,2,3,4,5,6,7,8,8}, []int{0,0,0,0,0,0,9,9,9,0}}, args{1, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := quickUnion{
				Items: tt.fields.Items,
				subtreeSize: tt.fields.subtreeSize,
			}
			q.Union(tt.args.i, tt.args.j)
			if !q.Connected(tt.args.i, tt.args.j) {
				t.Errorf("quickUnion.Union() = %v, %v not connected", tt.args.i, tt.args.j)
			}
		})
	}
}

func Test_quickUnion_Connected(t *testing.T) {
	type fields struct {
		Items []int
	}
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := quickUnion{
				Items: tt.fields.Items,
			}
			if got := q.Connected(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("quickUnion.Connected() = %v, want %v", got, tt.want)
			}
		})
	}
}
