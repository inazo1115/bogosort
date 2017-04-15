package bogosort

import (
	"math/rand"
	"sort"
)

func shuffle(xs sort.Interface) {
	for i := 0; i < xs.Len(); i++ {
		xs.Swap(i, rand.Intn(xs.Len()))
	}
}

func Bogosort(xs sort.Interface) {
	for {
		if sort.IsSorted(xs) {
			return
		}
		shuffle(xs)
	}
}

func BogosortWithCount(xs sort.Interface) int {
	i := 0
	for {
		if sort.IsSorted(xs) {
			return i
		}
		shuffle(xs)
		i++
	}
}
