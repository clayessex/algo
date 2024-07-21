package algo

import (
	"slices"
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

func TestRotate(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6}
	s = Rotate(s, 3)
	expect(t, s, []int{4, 5, 6, 1, 2, 3})
}

func TestCount(t *testing.T) {
	s := []int{1, 2, 5, 5, 6, 5, 9, 8, 5}
	i := Count(s, 5)
	expect(t, i, 4)
}

func TestMerge(t *testing.T) {
	data := []struct {
		name string
		a    []int
		b    []int
		want []int
	}{
		{"1", []int{}, []int{}, []int{}},
		{"2", []int{1, 2, 3}, []int{}, []int{1, 2, 3}},
		{"3", []int{}, []int{1, 2, 3}, []int{1, 2, 3}},
		{"4", []int{1, 2, 5, 7}, []int{3, 4, 6, 8, 9}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{"5", []int{1, 2, 6, 7}, []int{3, 4}, []int{1, 2, 3, 4, 6, 7}},
		{"6", []int{1, 2}, []int{3, 4}, []int{1, 2, 3, 4}},
		{"7", []int{3, 4}, []int{1, 2}, []int{1, 2, 3, 4}},
	}

	for _, v := range data {
		expect(t, Merge(v.a, v.b), v.want)
	}
}

func TestClamp(t *testing.T) {
	s := []int{1, 3, 5, 6, 8, 9}
	r := make([]int, 0, len(s))
	for _, v := range s {
		r = append(r, Clamp(v, 3, 6))
	}
	expect(t, r, []int{3, 3, 5, 6, 6, 6})
}

func TestMapKeys(t *testing.T) {
	s := make(map[int]int)
	s[3], s[9], s[7] = 1, 2, 3
	keys := MapKeys(s)
	slices.Sort(keys)
	expect(t, keys, []int{3, 7, 9})
}

func TestMapValues(t *testing.T) {
	s := make(map[int]int)
	s[3], s[9], s[7] = 1, 2, 3
	values := MapValues(s)
	slices.Sort(values)
	expect(t, values, []int{1, 2, 3})
}
