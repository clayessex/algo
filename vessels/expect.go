package vessels

import (
	"reflect"
	"testing"
)

func expect(t *testing.T, actual interface{}, expected interface{}) {
	t.Helper()
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("%s failed:\n    expected: %v\n      actual: %v\n", t.Name(), expected, actual)
	}
}

func expectNot(t *testing.T, actual interface{}, expected interface{}) {
	t.Helper()
	if reflect.DeepEqual(expected, actual) {
		t.Fatalf("%s failed:\n    expected: %v\n      actual: %v\n", t.Name(), expected, actual)
	}
}

func expectNil(t *testing.T, actual interface{}) {
	t.Helper()
	if !reflect.ValueOf(actual).IsNil() {
		t.Fatalf("%s failed: expected %v to be nil", t.Name(), actual)
	}
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
