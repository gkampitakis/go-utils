package arrays

func ReverseArrayString(input []string, modify bool) (reverse []string) {

	if modify {
		reverse = input
	} else {
		reverse = make([]string, len(input))
	}

	for i, j := 0, len(input)-1; i < len(input)/2; i, j = i+1, j+1 {

		reverse[i], reverse[j] = reverse[j], reverse[i]

	}

	return

}

func ReverseArrayInt(input []int, modify bool) (reverse []int) {

	if modify {
		reverse = input
	} else {
		reverse = make([]int, len(input))
	}

	for i, j := 0, len(input)-1; i < len(input)/2; i, j = i+1, j+1 {

		reverse[i], reverse[j] = reverse[j], reverse[i]

	}

	return

}
