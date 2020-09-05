package strings

import "testing"

func TestReverseStrings(t *testing.T) {
	testValues := []string{"Test", "civic", "go", "Golang"}
	wantValues := []string{"tseT", "civic", "og", "gnaloG"}

	t.Run("Reverse values", func(t *testing.T) {
		for i, test := range testValues {
			if got := ReverseString(test); got != wantValues[i] {
				t.Errorf("Wanted %v but got %v", wantValues[i], got)
			}
		}
	})

	t.Run("Empty Values", func(t *testing.T) {
		if got := ReverseString(""); got != "" {
			t.Errorf("Wanted %v but got %v", "", got)
		}
	})
}
