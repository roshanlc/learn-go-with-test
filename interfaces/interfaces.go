package interfaces

import "math"

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Length float64
	Width  float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (rec Rectangle) Area() float64 {
	return rec.Width * rec.Length
}
