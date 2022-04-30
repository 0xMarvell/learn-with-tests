package smi

import "math"

type Rectangle struct {
	length, width float64
}
type Circle struct {
	radius float64
}

func Perimeter(rec Rectangle) float64 {
	return 2 * (rec.length + rec.width)
}

func (rec Rectangle) Area() float64 {
	return rec.length * rec.width
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}
