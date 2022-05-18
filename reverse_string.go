package reverse_string

func ReverseString(input string) (output string) {
	for _, i := range input {
		output = string(i) + output
	}
	return output
}
