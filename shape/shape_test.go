package shape

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	input := Rectangle{10.0, 10.0}
	got := Perimeter(input)
	want := 40.0
	if got != want {
		t.Errorf("Perimeter(%v) = %v; want %v", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, got, want)
	}
}
