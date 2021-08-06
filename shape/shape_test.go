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

	checkArea := func(t *testing.T, name string, shape Shape, want float64) {
		t.Helper()

		got := shape.Area()

		if got != want {
			t.Errorf("Shape: %v; got = %v; want %v", name, got, want)
		}
	}

	var testcase_table = []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: &Rectangle{Width: 10.0, Height: 10.0}, want: 100.0},
		{name: "Rectangle", shape: &Rectangle{Width: 15.0, Height: 10.0}, want: 150.0},
		{name: "Circle", shape: &Circle{Radius: 10.0}, want: 314.1592653589793},
		{name: "Circle", shape: &Circle{Radius: 100.0}, want: 31415.926535897932},
		{name: "Triangle", shape: &Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	for _, v := range testcase_table {
		t.Run(v.name, func(t *testing.T) {
			name := v.name
			shape := v.shape
			want := v.want
			checkArea(t, name, shape, want)
		})
	}
}
