package algorithms

type QuickUnion interface {
	Union(i, j int)
	Connected(i, j int) bool
}

type quickUnion struct {
	Items []int
	subtreeSize []int
}

func NewQuickUnion(size int) QuickUnion {
	items := make([]int, size)
	subtreeSize := make([]int, size)

	for i := 0 ; i < size ; i++ {
		items[i] = i
		subtreeSize[i] = 1
	}

	return &quickUnion{items, subtreeSize}
}

// finds the root of the index node
func (q quickUnion) root(i int) int {
	for i != q.Items[i] {
		i = q.Items[i]
	}
	return i
}

// connects the elements at index j and i
func (q quickUnion) Union(i, j int) {
	x := q.root(i)
	y := q.root(j)

	if x == y {
		return
	}

	if q.subtreeSize[x] < q.subtreeSize[y] {
		q.Items[x] = y
		q.subtreeSize[y] += q.subtreeSize[x]
	} else {
		q.Items[y] = x
		q.subtreeSize[x] += q.subtreeSize[y]
	}

}

// find whether the elements at i and j are connected or not
func (q quickUnion) Connected(i, j int) bool {
	return q.root(i) == q.root(j)
}
