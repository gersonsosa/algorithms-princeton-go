package intersection

type Point struct {
	X, Y float64
}

func Intersect(a, b []Point) int {

	count := 0
	append(a, b...)
	// shell sort a
	for i := 1; i < len(a); i++ {
		if a[i-1] == a[i] {
			count++
		}
	}
	return count
}
