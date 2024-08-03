package expected

import (
	"reflect"
	"testing"
)

type TestAdapter struct {
	t     *testing.T
	value interface{}
}

func New(t *testing.T) *TestAdapter {
	adapt := new(TestAdapter)
	adapt.t = t
	return adapt
}

func (adapt *TestAdapter) Expect(value interface{}) *TestAdapter {
	adapt.t.Helper()
	adapt.value = value
	return adapt
}

func (adapt *TestAdapter) ExpectOk(value interface{}, ok bool) *TestAdapter {
	adapt.t.Helper()
	if !ok {
		adapt.t.Fatalf("%s failed: expected ok", adapt.t.Name())
	}
	adapt.value = value
	return adapt
}

func (adapt *TestAdapter) ToBe(want interface{}) {
	adapt.t.Helper()
	if !reflect.DeepEqual(want, adapt.value) {
		adapt.t.Fatalf("%s failed:\n    expected: %v\n      actual: %v\n", adapt.t.Name(), want, adapt.value)
	}
}

func (adapt *TestAdapter) ToNotBe(want interface{}) {
	adapt.t.Helper()
	if reflect.DeepEqual(want, adapt.value) {
		adapt.t.Fatalf("%s failed:\n  not expected: %v\n        actual: %v\n", adapt.t.Name(), want, adapt.value)
	}
}

func (adapt *TestAdapter) Assert(actual interface{}, want interface{}) {
	adapt.t.Helper()
	if !reflect.DeepEqual(want, actual) {
		adapt.t.Fatalf("%s failed:\n    expected: %v\n      actual: %v\n", adapt.t.Name(), want, actual)
	}
}
