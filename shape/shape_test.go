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
	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Area()
		want := 100.0

		if got != want {
			t.Errorf("got = %v; want %v", got, want)
		}
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{10.0}
		got := circle.Area()
		want := 314.16

		if got != want {
			t.Errorf("got = %v; want %v", got, want)
		}
	})
}
