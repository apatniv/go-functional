package iter_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/BooleanCat/go-functional/internal/assert"
	"github.com/BooleanCat/go-functional/iter"
	"github.com/BooleanCat/go-functional/iter/ops"
)

func ExampleFold() {
	sum := iter.Fold[int](iter.Take[int](iter.Count(), 4), 0, ops.Add[int])

	fmt.Println(sum)
	// Output: 6
}

func TestCollect(t *testing.T) {
	items := iter.Collect[int](iter.Take[int](iter.Count(), 5))
	assert.SliceEqual(t, items, []int{0, 1, 2, 3, 4})
}

func TestCollectEmpty(t *testing.T) {
	items := iter.Collect[int](iter.Take[int](iter.Count(), 0))
	assert.Empty(t, items)
}

func TestFold(t *testing.T) {
	add := func(a, b int) int { return a + b }
	sum := iter.Fold[int](iter.Take[int](iter.Count(), 11), 0, add)
	assert.Equal(t, sum, 55)

	concat := func(path string, part int) string {
		return path + strconv.Itoa(part) + "/"
	}
	result := iter.Fold[int](iter.Take[int](iter.Count(), 3), "/", concat)
	assert.Equal(t, result, "/0/1/2/")
}