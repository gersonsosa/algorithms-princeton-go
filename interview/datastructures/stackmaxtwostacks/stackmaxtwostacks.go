package stackmaxtwostacks

type (
	StackMaxTwoStacks struct {
		top *node
		length int
	}
	node struct {
		value interface{}
		prev *node
	}
)

func New() *StackMaxTwoStacks {
	return &StackMaxTwoStacks{nil, 0}
}

func (stack *StackMaxTwoStacks) Push(value interface{}) {
	stack.top = &node{value, stack.top}
	stack.length++
}

func (stack *StackMaxTwoStacks) Pop() interface{} {
	stack.length--
	oldTop := stack.top
	stack.top = stack.top.prev
	return oldTop
}