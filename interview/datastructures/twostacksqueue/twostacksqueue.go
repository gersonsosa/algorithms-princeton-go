package twostacksqueue

import "github.com/golang-collections/collections/stack"

type TwoStackQueue struct {
	Stack1 *stack.Stack
	Stack2 *stack.Stack
}

func (twoStackQueue TwoStackQueue) Queue(i int) {
	twoStackQueue.Stack1.Push(i)
}

func (twoStackQueue TwoStackQueue) Dequeue() int {
	if twoStackQueue.Stack2.Len() == 0 {
		for twoStackQueue.Stack1.Len() > 0 {
			twoStackQueue.Stack2.Push(twoStackQueue.Stack1.Pop())
		}
	}

	return twoStackQueue.Stack2.Pop().(int)
}

func (twoStackQueue TwoStackQueue) len() int {
	return twoStackQueue.Stack2.Len() + twoStackQueue.Stack1.Len()
}