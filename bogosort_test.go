package bogosort

import (
	"fmt"
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

	actual = Ints([]int{0, 3, 7, 2, 6, 8, 9, 1, 4, 5})
	Bogosort(actual)
	if !actual.Eq(expected) {
		t.Errorf("error")
	}

	actual = Ints([]int{3, 8, 5, 9, 0, 4, 1, 2, 6, 7})
	i := BogosortWithCount(actual)
	if !actual.Eq(expected) {
		t.Errorf("error")
	}
	fmt.Printf("count: %d\n", i)
}
