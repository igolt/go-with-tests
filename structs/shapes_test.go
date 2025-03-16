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
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "rectangle area", shape: Rectangle{10, 10}, hasArea: 100},
		{name: "circle area", shape: Circle{10}, hasArea: 314.1592653589793},
		{name: "triangle area", shape: Triangle{12, 6}, hasArea: 36},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g expected %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}
