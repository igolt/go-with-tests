package structs

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(Rectangle{10, 10})
	var expected float64 = 40

	if got != expected {
		t.Errorf("got %.2f expected %.2f", got, expected)
	}
}

func TestArea(t *testing.T) {
	t.Run("rectangle area", func(t *testing.T) {
		rectangle := Rectangle{10, 10}

		got := rectangle.Area()
		var expected float64 = 100

		if got != expected {
			t.Errorf("got %.2f expected %.2f", got, expected)
		}
	})

	t.Run("circle area", func(t *testing.T) {
		circle := Circle{10}

		got := circle.Area()
		var expected float64 = 314.1592653589793

		if got != expected {
			t.Errorf("got %g expected %g", got, expected)
		}
	})
}
