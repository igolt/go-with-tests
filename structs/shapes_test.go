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
	areaTests := []struct {
		shape    Shape
		expected float64
	}{
		{Rectangle{10, 10}, 100},
		{Circle{10}, 314.1592653589793},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.expected {
			t.Errorf("got %g expected %g", got, tt.expected)
		}
	}
}
