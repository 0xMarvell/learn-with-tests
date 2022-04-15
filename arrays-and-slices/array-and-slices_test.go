package arraysandslices

import "testing"

func TestSum(t *testing.T) {
	testcases := []struct {
		numbers [5]int
		sum     int
	}{
		{[5]int{1, 2, 3, 4, 5}, 15},
		{[5]int{-3, 4, 1, 13, 42}, 57},
		{[5]int{9, 20, 0, 0, 5}, 34},
		{[5]int{1, -9, 71, -6, -3}, 54},
	}

	for _, test := range testcases {
		s := Sum(test.numbers)
		if s != test.sum {
			t.Errorf("Sum(%d): %d. Expected: %d, Got: %d", test.numbers, s, test.sum, s)
		}
	}
}
