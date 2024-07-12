package vessels

import (
	"testing"

	"github.com/clayessex/godev/expected"
)

func expect[T any](t *testing.T, actual T, want T) {
	t.Helper()
	expected.Expect(t, actual, want)
}

func expectNot(t *testing.T, actual interface{}, want interface{}) {
	t.Helper()
	expected.ExpectNot(t, actual, want)
}

func expectNil(t *testing.T, actual interface{}) {
	t.Helper()
	expected.ExpectNil(t, actual)
}

// List specific
func expectSequence(t *testing.T, first *ListNode[int], last *ListNode[int]) {
	t.Helper()
	count := first.value
	for i := first; i != last; i = i.next {
		expect(t, i.value, count)
		count++
	}
}
