package iteration

func Repeat(char string, times int) string {
	var output string
	for i := 0; i < times; i++ {
		output += char
	}
	return output
}
