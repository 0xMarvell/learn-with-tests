package iteration

func Repeat(char string, reps int) string {
	if char == "" {
		return "Invalid input"
	}

	res := ""
	for i := 0; i < reps; i++ {
		res += char
	}
	return res
}
