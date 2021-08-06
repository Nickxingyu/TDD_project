package shape

type Rectangle struct {
	Width  float64
	Height float64
}

func Perimeter(shape Rectangle) float64 {
	return 2 * (shape.Width + shape.Height)
}
