package asserts

import "testing"

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	AssertEqual(t, got, true)
}

func AssertFalse(t *testing.T, got bool) {
	t.Helper()
	AssertEqual(t, got, false)
}

func AssertEqual[T comparable](t *testing.T, got, expected T) {
	t.Helper()
	if expected != got {
		t.Errorf("expected %#v but got %#v", expected, got)
	}
}

func AssertNotEqual[T comparable](t *testing.T, got, expected T) {
	t.Helper()
	if expected == got {
		t.Errorf("expected %v to be not equal to %v", expected, got)
	}
}
