package generics_test

import (
	"testing"

	. "github.com/igolt/go-with-tests/asserts"
	"github.com/igolt/go-with-tests/generics"
)

type (
	StackInt    = generics.Stack[int]
	StackString = generics.Stack[string]
)

func TestStack(t *testing.T) {
	t.Run("test stack of ints", func(t *testing.T) {
		stack := StackInt{}

		stack.Push(2)
		stack.Push(4)
		stack.Push(3)

		value, ok := stack.Pop()

		AssertEqual(t, value, 3)
		AssertTrue(t, ok)

		value, ok = stack.Pop()

		AssertEqual(t, value, 4)
		AssertTrue(t, ok)

		value, ok = stack.Pop()

		AssertEqual(t, value, 2)
		AssertTrue(t, ok)

		value, ok = stack.Pop()

		AssertEqual(t, value, 0)
		AssertFalse(t, ok)
	})

	t.Run("test stack of strings", func(t *testing.T) {
		stack := StackString{}

		stack.Push("Go")
		stack.Push("Lang")
		stack.Push("Rocks")

		value, ok := stack.Pop()

		AssertEqual(t, value, "Rocks")
		AssertTrue(t, ok)

		value, ok = stack.Pop()

		AssertEqual(t, value, "Lang")
		AssertTrue(t, ok)

		value, ok = stack.Pop()

		AssertEqual(t, value, "Go")
		AssertTrue(t, ok)

		value, ok = stack.Pop()

		AssertEqual(t, value, "")
		AssertFalse(t, ok)
	})
}
