package shape

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Perimeter()
		want := 40.0

		if got != want {
			t.Errorf("got = %v; want %v", got, want)
		}
	})
}

func TestArea(t *testing.T) {

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()

		got := shape.Area()

		if got != want {
			t.Errorf("got = %v; want %v", got, want)
		}
	}

	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		want := 100.0
		checkArea(t, rectangle, want)
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{10.0}
		want := 314.1592653589793
		checkArea(t, circle, want)
	})
}
