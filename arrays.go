package arrays

import q3 "rsc.io/quote/v3"

func reverseArrayString(input []string) (reverse []string) {

	reverse = make([]string, len(input))

	for i, element := range input {

		reverse[len(input)-1-i] = element

	}

	return

}

func Proverb() string {
	return q3.Concurrency()
}

func HelloWorld() string {
	return q3.HelloV3()
}
