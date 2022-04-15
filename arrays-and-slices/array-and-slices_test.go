package arraysandslices

import "testing"

func TestSumArr(t *testing.T) {
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
		s := SumArr(test.numbers)
		if s != test.sum {
			t.Errorf("Sum(%d): %d. Expected: %d, Got: %d", test.numbers, s, test.sum, s)
		}
	}
}

func TestSumSlice(t *testing.T) {
	testcases := []struct {
		numbers []int
		sum     int
	}{
		{[]int{1, 2, 3, 4}, 10},
		{[]int{-3, 4, 1, 13, 42, -33}, 24},
		{[]int{9, 20, 0, 0, 5, -6, -54}, -26},
		{[]int{1, -9, 71, -6, -3}, 54},
	}

	for _, test := range testcases {
		s := SumSlice(test.numbers)
		if s != test.sum {
			t.Errorf("Sum(%d): %d. Expected: %d, Got: %d", test.numbers, s, test.sum, s)
		}
	}
}
