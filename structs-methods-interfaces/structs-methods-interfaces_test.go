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
	t.Run("testing rectangles", func(t *testing.T) {
		testcases := []struct {
			rectangle Rectangle
			area      float64
		}{
			{Rectangle{10, 10}, 100},
			{Rectangle{3.5, 6.72}, 23.52},
			{Rectangle{13, 77.95}, 1013.35},
		}

		for _, test := range testcases {
			a := test.rectangle.Area()
			if a != test.area {
				t.Errorf("Area(%.2f): %.2f. Expected: %.2f", test.rectangle, a, test.area)
			}
		}
	})

	// t.Run("testing circles", func(t *testing.T) {
	// 	testcases := []struct {
	// 		circle Circle
	// 		area   float64
	// 	}{
	// 		{Circle{10}, 314.159265},
	// 		{Circle{3.5}, 38.484510},
	// 		{Circle{13}, 530.929158},
	// 	}

	// 	for _, test := range testcases {
	// 		a := test.circle.Area()
	// 		if a != test.area {
	// 			fmt.Printf("a=%f, test.area=%f\n", a, test.area)
	// 			fmt.Println(test.area == a)
	// 			t.Errorf("Area(%f): %f. Expected: %f", test.circle, a, test.area)
	// 		}
	// 	}
	// })
	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()

		want := 314.1592653589793
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})

}
