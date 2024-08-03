package expected

import (
	"reflect"
	"testing"
)

type TestAdapter testing.T

type TestAdapterValue struct {
	t     *TestAdapter
	value interface{}
}

func New(t *testing.T) *TestAdapter {
	return (*TestAdapter)(t)
}

func (adapt *TestAdapter) Expect(value interface{}) *TestAdapterValue {
	adapt.Helper()
	t := &TestAdapterValue{adapt, value}
	return t
}

func (adapt *TestAdapter) ExpectOk(value interface{}, ok bool) *TestAdapterValue {
	adapt.Helper()
	if !ok {
		adapt.Fatalf("%s failed: expected ok", adapt.Name())
	}
	t := &TestAdapterValue{adapt, value}
	return t
}

func (adapt *TestAdapter) ExpectNotOk(_ interface{}, ok bool) {
	adapt.Helper()
	if ok {
		adapt.Fatalf("%s failed: expected not ok", adapt.Name())
	}
}

func (v *TestAdapterValue) ToBe(want interface{}) {
	v.t.Helper()
	if !reflect.DeepEqual(want, v.value) {
		v.t.Fatalf("%s failed:\n    expected: %v\n      actual: %v\n", v.t.Name(), want, v.value)
	}
}

func (v *TestAdapterValue) ToNotBe(want interface{}) {
	v.t.Helper()
	if reflect.DeepEqual(want, v.value) {
		v.t.Fatalf("%s failed:\n  not expected: %v\n        actual: %v\n", v.t.Name(), want, v.value)
	}
}

func (adapt *TestAdapter) Assert(actual interface{}, want interface{}) {
	adapt.Helper()
	if !reflect.DeepEqual(want, actual) {
		adapt.Fatalf("%s failed:\n    expected: %v\n      actual: %v\n", adapt.Name(), want, actual)
	}
}
