package vessels

import (
	"cmp"
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

func logListNodes(t *testing.T, first *ListNode[int], last *ListNode[int]) {
	t.Helper()
	for ; first != last; first = first.next {
		t.Logf("%v\n", first.value)
	}
}

func compareNodes[T cmp.Ordered](t *testing.T,
	first1 *ListNode[T], last1 *ListNode[T],
	first2 *ListNode[T], last2 *ListNode[T],
) {
	t.Helper()

	for first1 != last1 && first2 != last2 {
		t.Logf("%v = %v\n", first1.value, first2.value)
		if first1.value != first2.value {
			t.Fatalf("%s failed:\n    expected: %v\n      actual: %v\n",
				t.Name(), first1.value, first2.value)
		}
		first1, first2 = first1.next, first2.next
	}

	if first1 != last1 || first2 != last2 {
		t.Fatalf("%s failed: array sizes unequal", t.Name())
	}
}

func compareNodesSlice[T cmp.Ordered](t *testing.T,
	first *ListNode[T], last *ListNode[T],
	want []T,
) {
	t.Helper()

	i := 0
	for first != last && i < len(want) {
		if first.value != want[i] {
			t.Fatalf("%s failed:\n    expected: %v\n      actual: %v\n",
				t.Name(), first.value, want[i])
		}
		first, i = first.next, i+1
	}
}

func expectOrdered[T cmp.Ordered](t *testing.T,
	first *ListNode[T], last *ListNode[T],
) {
	//
	t.Helper()

	lastvalue := first.value
	first = first.next

	for first != last {
		if !(lastvalue <= first.value) {
			t.Fatalf("%s failed - unordered:\n  prev: %v\n  next: %v\n",
				t.Name(), lastvalue, first.value)
		}
		first = first.next
	}
}
