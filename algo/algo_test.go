package algo

import (
	"testing"

	"github.com/clayessex/godev/expected"
)

func expect[T any](t *testing.T, actual T, want T) {
	t.Helper()
	expected.Expect(t, actual, want)
}

func TestMap(t *testing.T) {
	s := []int{1, 2, 3, 4}
	add2 := Map(s, func(v int) float64 { return float64(v) + 2.0 })
	expect(t, add2, []float64{3.0, 4.0, 5.0, 6.0})
	mul2 := Map(s, func(v int) int { return v * 2 })
	expect(t, mul2, []int{2, 4, 6, 8})
}

func TestReduce(t *testing.T) {
	s := []int{1, 2, 3, 4}
	sum := Reduce(s, 0, func(a int, v int) int { return a + v })
	expect(t, sum, 10)
	mul := Reduce(s, 1.0, func(a float64, v int) float64 { return a * float64(v) })
	expect(t, mul, 24.0)
}

func TestFilter(t *testing.T) {
	s := []int{1, 2, 3, 4}
	f := Filter(s, func(v int) bool { return v != 3 })
	expect(t, f, []int{1, 2, 4})
}
