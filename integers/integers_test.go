package integers

import "testing"

func TestAdder(t *testing.T) {
	testcases := []struct{ v1, v2, sum float64 }{
		{1, 2, 3},
		{-2, 3, 1},
		{3.142, 5.663, 8.805},
		{-44.2112, -108.33, -152.5412},
		{-3, 0, -3},
		{0, 0, 0},
	}

	for _, tt := range testcases {
		s := Adder(tt.v1, tt.v2)
		if s != tt.sum {
			t.Errorf("Adder(%f, %f) returned %f. Want %f", tt.v1, tt.v2, s, tt.sum)
		}
	}
}
