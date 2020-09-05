package strings

func ReverseString(input string) string {

	return string(reverseRune([]rune(input)))

}

func reverseRune(input []rune) []rune {

	for i, j := 0, len(input)-1; i < len(input)/2; i, j = i+1, j-1 {

		input[i], input[j] = input[j], input[i]

	}

	return input

}
