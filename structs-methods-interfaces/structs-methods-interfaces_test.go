package smi

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	testcases := []struct {
		rectangle Rectangle
		perimeter float64
	}{
		{Rectangle{10, 10}, 40},
		{Rectangle{0.5, 3.442}, 7.884},
		{Rectangle{12, 444.24}, 912.48},
	}

	for _, test := range testcases {
		p := Perimeter(test.rectangle)
		if p != test.perimeter {
			t.Errorf("Perimeter(%.3f): %.3f. Expected: %.3f", test.rectangle, p, test.perimeter)
		}
	}
}

func TestArea(t *testing.T) {
	testcases := []struct {
		shape Shape
		area  float64
	}{
		{Rectangle{10, 10}, 100},
		{Rectangle{3.5, 6.72}, 23.52},
		{Rectangle{13, 77.95}, 1013.35},
		{Circle{10}, 314.1592653589793},
		{Circle{3.5}, 38.48451000647496},
		{Circle{13}, 530.929158456675},
		{Triangle{12, 6}, 36.0},
	}

	for _, test := range testcases {
		a := test.shape.Area()
		if a != test.area {
			t.Errorf("Area(%g): %g. Expected: %g", test.shape, a, test.area)
		}
	}
}
