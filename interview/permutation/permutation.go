package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("check unit test")
}

func IsPermutation(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	sort.Ints(a)
	sort.Ints(b)

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
