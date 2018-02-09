package stackmax

type StackMax struct {
	Items []int
	max int
	last int
}

func New() *StackMax {
	return &StackMax{make([]int, 0, 10), -1, -1}
}

func (stackMax StackMax) findMax() int {
	return stackMax.Items[stackMax.max]
}

func (stackMax *StackMax) Push(i int) {
	stackMax.Items = append(stackMax.Items, i)
	stackMax.last += 1

	if i > stackMax.Items[stackMax.max] {
		stackMax.max = stackMax.last
	}
}

func (stackMax *StackMax) Pop() int {
	if stackMax.last == stackMax.max {
		defer stackMax.searchNewMax()
	}

	defer func() {
		stackMax.Items = stackMax.Items[:stackMax.last]
		stackMax.last--
		stackMax.shrinkIfNecessary()
	}()

	popped := stackMax.Items[stackMax.last]
	return popped
}

func (stackMax *StackMax) searchNewMax() {
	stackMax.max = 0
	for i := 1; i < len(stackMax.Items); i++ {
		if stackMax.Items[i] > stackMax.Items[stackMax.max] {
			stackMax.max = i
		}
	}
}

func (stackMax *StackMax) shrinkIfNecessary() {
	n := len(stackMax.Items)
	if cap(stackMax.Items) / 4 >= n {
		newSlice := make([]int, (n+1)*2)
		copy(newSlice, stackMax.Items)
		stackMax.Items = newSlice
	}
}
