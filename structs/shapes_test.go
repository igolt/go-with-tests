package structs

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(10, 10)
	var expected float64 = 40

	if got != expected {
		t.Errorf("got %.2f expected %.2f", got, expected)
	}
}

func TestArea(t *testing.T) {
	got := Area(10, 10)
	var expected float64 = 100

	if got != expected {
		t.Errorf("got %.2f expected %.2f", got, expected)
	}
}
