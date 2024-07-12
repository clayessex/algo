package expected

import (
	"reflect"
	"testing"
)

func Expect(t *testing.T, actual interface{}, expected interface{}) {
	t.Helper()
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("%s failed:\n    expected: %v\n      actual: %v\n", t.Name(), expected, actual)
	}
}

func ExpectNot(t *testing.T, actual interface{}, expected interface{}) {
	t.Helper()
	if reflect.DeepEqual(expected, actual) {
		t.Fatalf("%s failed:\n    expected: %v\n      actual: %v\n", t.Name(), expected, actual)
	}
}

func ExpectNil(t *testing.T, actual interface{}) {
	t.Helper()
	if !reflect.ValueOf(actual).IsNil() {
		t.Fatalf("%s failed: expected %v to be nil", t.Name(), actual)
	}
}
