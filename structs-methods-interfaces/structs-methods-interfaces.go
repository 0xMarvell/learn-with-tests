package smi

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	length, width float64
}
type Circle struct {
	radius float64
}
type Triangle struct {
	base, height float64
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

func (t Triangle) Area() float64 {
	return 0.5 * t.base * t.height
}
