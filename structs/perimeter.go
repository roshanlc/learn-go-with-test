package structs

type Rectangle struct {
	length float64
	width  float64
}

// Perimeter returns the perimeter of a rectangle
func Perimeter(rec Rectangle) float64 {
	return 2 * (rec.length + rec.width)
}

// Area returns the area of a rectangle
func Area(rec Rectangle) float64 {
	return rec.length * rec.width
}
