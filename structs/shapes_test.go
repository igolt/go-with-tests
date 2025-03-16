package structs

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(Rectangle{10, 10})
	var expected float64 = 40

	if got != expected {
		t.Errorf("got %.2f expected %.2f", got, expected)
	}
}

type Shape interface {
	Area() float64
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, expected float64) {
		t.Helper()
		got := shape.Area()

		if got != expected {
			t.Errorf("got %g expected %g", got, expected)
		}
	}

	t.Run("rectangle area", func(t *testing.T) {
		checkArea(t, Rectangle{10, 10}, 100)
	})

	t.Run("circle area", func(t *testing.T) {
		checkArea(t, Circle{10}, 314.1592653589793)
	})
}
