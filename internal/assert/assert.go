package assert

import (
	"testing"
)

func Equal[T comparable](t *testing.T, actual, expected T) {
	t.Helper()
	if actual != expected {
		t.Errorf("got: %v; want: %v", actual, expected)
	}
}

func EqualSlice[E comparable](t *testing.T, actual, expected []E) {
	if len(actual) != len(expected) {
		t.Errorf("got: %v; want: %v", actual, expected)
	}
	for i := range actual {
		if actual[i] != expected[i] {
			t.Errorf("got: %v; want: %v", actual, expected)
		}
	}
}
