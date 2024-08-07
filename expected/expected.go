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

type Expected testing.T

type ExpectedValue struct {
	t     *Expected
	value interface{}
}

func New(t *testing.T) *Expected {
	return (*Expected)(t)
}

func (x *Expected) Expect(value interface{}) *ExpectedValue {
	x.Helper()
	t := &ExpectedValue{x, value}
	return t
}

func (x *Expected) ExpectOk(value interface{}, ok bool) *ExpectedValue {
	x.Helper()
	if !ok {
		x.Fatalf("%s failed: expected ok", x.Name())
	}
	t := &ExpectedValue{x, value}
	return t
}

func (x *Expected) ExpectNotOk(_ interface{}, ok bool) {
	x.Helper()
	if ok {
		x.Fatalf("%s failed: expected not ok", x.Name())
	}
}

func (x *Expected) ExpectErr(_ interface{}, err error) *ExpectedValue {
	x.Helper()
	if err == nil {
		x.Fatalf("%s failed: expected error got nil", x.Name())
	}
	t := &ExpectedValue{x, err}
	return t
}

func (x *Expected) ExpectErrNil(value interface{}, err error) *ExpectedValue {
	x.Helper()
	if err != nil {
		x.Fatalf("%s failed: expected nil error got: %v", x.Name(), err)
	}
	t := &ExpectedValue{x, value}
	return t
}

func (v *ExpectedValue) ToBe(want interface{}) {
	v.t.Helper()
	if !reflect.DeepEqual(want, v.value) {
		v.t.Fatalf("%s failed:\n    expected: %v\n      actual: %v\n", v.t.Name(), want, v.value)
	}
}

func (v *ExpectedValue) ToNotBe(want interface{}) {
	v.t.Helper()
	if reflect.DeepEqual(want, v.value) {
		v.t.Fatalf("%s failed:\n  not expected: %v\n        actual: %v\n", v.t.Name(), want, v.value)
	}
}

func (v *ExpectedValue) Value() interface{} {
	return v.value
}

func (x *Expected) Assert(actual interface{}, want interface{}) {
	x.Helper()
	if !reflect.DeepEqual(want, actual) {
		x.Fatalf("%s failed:\n    expected: %v\n      actual: %v\n", x.Name(), want, actual)
	}
}
