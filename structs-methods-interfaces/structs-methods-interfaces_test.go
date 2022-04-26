package smi

import "testing"

func TestPerimeter(t *testing.T) {
	testcases := []struct{ length, width, perimeter float64 }{
		{10, 10, 40},
		{0.5, 3.442, 7.884},
		{12, 444.24, 912.48},
	}

	for _, test := range testcases {
		p := Perimeter(test.length, test.width)
		if p != test.perimeter {
			t.Errorf("Perimeter(%.3f, %.3f): %.3f. Expected: %.3f", test.length, test.width, p, test.perimeter)
		}
	}
}
