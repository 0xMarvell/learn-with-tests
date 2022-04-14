package iteration

const reps = 5

func Repeat(char string) string {
	if char == "" {
		return "Invalid input"
	}

	res := ""
	for i := 0; i < reps; i++ {
		res += char
	}
	return res
}
