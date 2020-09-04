package goUtils

func ReverseArrayString(input []string) (reverse []string) {

	reverse = make([]string, len(input))

	for i, element := range input {

		reverse[len(input)-1-i] = element

	}

	return

}
