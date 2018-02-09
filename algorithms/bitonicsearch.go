package algorithms

func Find2lgn(x int, a []int) int {
	pivot := 0 + (len(a) - 0) / 2

	found1 := find(x, a, 0, pivot)
	found2 := findInv(x, a, pivot, len(a))
	if found1 > -1 {
		return found1
	} else if found2 > -1 {
		return found2
	}
	return -1
}

func Find(x int, a []int) int {

	pivot := -1
	start := 0
	end := len(a)
	for pivot == -1 {
		if start == end {
			pivot = start
		}

		middle := start + (end - start) / 2

		if a[middle] > a[middle + 1] {
			end = middle
		} else {
			start = middle + 1
		}
	}
	found1 := find(x, a, 0, pivot)
	found2 := find(x, a, pivot, len(a))
	if found1 > -1 {
		return found1
	} else if found2 > -1 {
		return found2
	}
	return -1
}

func find(x int, a []int, start int, end int) int {
	for start < end {
		middle := start + (end - start) / 2

		if x < a[middle] {
			end = middle - 1
		} else if x > a[middle] {
			start = middle + 1
		} else {
			return middle
		}
	}
	return -1
}

func findInv(x int, a []int, start int, end int) int {
	for start < end {
		middle := start + (end - start) / 2

		if x > a[middle] {
			end = middle - 1
		} else if x < a[middle] {
			start = middle + 1
		} else {
			return middle
		}
	}
	return -1
}