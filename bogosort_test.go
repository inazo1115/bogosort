package bogosort

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
)

type Ints []int

func (xs Ints) Len() int {
	return len([]int(xs))
}

func (xs Ints) Less(i, j int) bool {
	return xs[i] < xs[j]
}

func (xs Ints) Swap(i, j int) {
	xs[i], xs[j] = xs[j], xs[i]
}

func (xs Ints) Eq(ys Ints) bool {
	for i := 0; i < xs.Len(); i++ {
		if xs[i] != ys[i] {
			return false
		}
	}
	return true
}

func init() {
	rand.Seed(1)
}

func TestBogosort(t *testing.T) {
	var actual Ints
	expected := Ints([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	// Normal
	actual = Ints([]int{0, 3, 7, 2, 6, 8, 9, 1, 4, 5})
	i, err := Bogosort(actual, math.MaxInt64)
	if err != nil {
		t.Errorf("...")
	}
	if !actual.Eq(expected) {
		t.Errorf("...")
	}
	fmt.Printf("count: %d\n", i)

	// Abnormal
	actual = Ints([]int{0, 3, 7, 2, 6, 8, 9, 1, 4, 5})
	if _, err := Bogosort(actual, 1); err == nil {
		t.Errorf("...")
	}
}
