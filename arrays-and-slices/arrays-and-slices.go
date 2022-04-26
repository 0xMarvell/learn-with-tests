package arraysandslices

func SumArr(numbers [5]int) int {
	total := 0
	for _, v := range numbers {
		total += v
	}
	return total

}

func SumSlice(numbers []int) int {
	total := 0
	for _, v := range numbers {
		total += v
	}
	return total

}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, SumSlice(numbers))
	}
	return sums
}
