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
	return []int{0}
}
