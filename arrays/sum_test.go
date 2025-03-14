package arrays

import (
	"fmt"
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		expected := 15

		if got != expected {
			t.Errorf("got %d expected %d given, %v", got, expected, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		expected := 6

		if got != expected {
			t.Errorf("got %d expected %d given, %v", got, expected, numbers)
		}
	})
}

func ExampleSum() {
	sum := Sum([]int{10, 20, 30, 40})
	fmt.Println(sum)
	// Output: 100
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	expected := []int{3, 9}

	if !slices.Equal(got, expected) {
		t.Errorf("got %d expected %d", got, expected)
	}
}

func ExampleSumAll() {
	sums := SumAll([]int{10, 20}, []int{1, 2, 3})
	fmt.Println(sums)
	// Output: [30 6]
}

func TestSumAllTails(t *testing.T) {
	t.Run("sum tails of slices of 2 elements", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}

		if !slices.Equal(got, expected) {
			t.Errorf("got %v expected %v", got, expected)
		}
	})

	t.Run("sum tails of slices with different numbers of elements", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3, 4, 5}, []int{0, 1, 9})
		expected := []int{14, 10}

		if !slices.Equal(got, expected) {
			t.Errorf("got %v expected %v", got, expected)
		}
	})

	t.Run("sum tails of slices with one element", func(t *testing.T) {
		got := SumAllTails([]int{1}, []int{9})
		expected := []int{0, 0}

		if !slices.Equal(got, expected) {
			t.Errorf("got %v expected %v", got, expected)
		}
	})

	t.Run("sum tail of an empty slice", func(t *testing.T) {
		got := SumAllTails([]int{})
		expected := []int{0}

		if !slices.Equal(got, expected) {
			t.Errorf("got %v expected %v", got, expected)
		}
	})
}

func ExampleSumAllTails() {
	sums := SumAllTails([]int{10, 20}, []int{1, 2, 3})
	fmt.Println(sums)
	// Output: [20 5]
}
