package shape

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	got := Perimeter([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	want := 55
	if got != want {
		t.Errorf("Perimeter(%v) = %v; want %v", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, got, want)
	}
}
