package twostacksqueue

import (
	"testing"

	"github.com/golang-collections/collections/stack"
)

func TestTwoStackQueue_Queue(t *testing.T) {
	type fields struct {
		Stack1 *stack.Stack
		Stack2 *stack.Stack
	}
	type args struct {
		i int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"should queue a element in to the first stack", fields{stack.New(), stack.New()}, args{0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			twoStackQueue := TwoStackQueue{
				Stack1: tt.fields.Stack1,
				Stack2: tt.fields.Stack2,
			}
			twoStackQueue.Queue(tt.args.i)
		})
	}
}

func TestTwoStackQueue_Dequeue(t *testing.T) {
	type fields struct {
		Stack1 *stack.Stack
		Stack2 *stack.Stack
	}
	i := stack.New()
	i.Push(0)

	j := stack.New()
	j.Push(1)
	j.Push(0)
	j.Push(3)
	j.Push(6)
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"should dequeue an element from the queue", fields{i, stack.New()}, 0},
		{"should dequeue first element in the queue", fields{j, stack.New()}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			twoStackQueue := TwoStackQueue{
				Stack1: tt.fields.Stack1,
				Stack2: tt.fields.Stack2,
			}
			if got := twoStackQueue.Dequeue(); got != tt.want {
				t.Errorf("TwoStackQueue.Dequeue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTwoStackQueue_len(t *testing.T) {
	type fields struct {
		Stack1 *stack.Stack
		Stack2 *stack.Stack
	}

	j := stack.New()
	j.Push(1)
	j.Push(0)
	j.Push(3)
	j.Push(6)

	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"should return the queue lenght", fields{j, stack.New()}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			twoStackQueue := TwoStackQueue{
				Stack1: tt.fields.Stack1,
				Stack2: tt.fields.Stack2,
			}
			if got := twoStackQueue.len(); got != tt.want {
				t.Errorf("TwoStackQueue.len() = %v, want %v", got, tt.want)
			}
		})
	}
}
