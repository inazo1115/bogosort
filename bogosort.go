package bogosort

import (
	"fmt"
	"math/rand"
	"sort"
)

func shuffle(xs sort.Interface) {
	for i := 0; i < xs.Len(); i++ {
		xs.Swap(i, rand.Intn(xs.Len()))
	}
}

func Bogosort(xs sort.Interface, limit uint) (uint, error) {
	var i uint = 0
	for {
		if sort.IsSorted(xs) {
			return i, nil
		}
		if i > limit {
			return 0, fmt.Errorf("try count is over limit: %d", limit)
		}
		shuffle(xs)
		i++
	}
}
