package arrays

import (
	"reflect"
	"testing"
)

func TestReverseArrayString(t *testing.T) {
	testValues := [][]string{[]string{"Hello World", "Test", "Blue"}, []string{"1", "2", "3"}}
	wantValues := [][]string{[]string{"Blue", "Test", "Hello World"}, []string{"3", "2", "1"}}

	t.Run("Reverse Arrays", func(t *testing.T) {
		for i, test := range testValues {
			if got := ReverseArrayString(test, true); !reflect.DeepEqual(got, wantValues[i]) {
				t.Errorf("Wanted %v but got %v", wantValues[i], got)
			}
		}
	})

	t.Run("Empty array", func(t *testing.T) {
		got := ReverseArrayString([]string{}, false)

		if len(got) > 0 {
			t.Errorf("Wanted [] but got %v", got)
		}
	})

	t.Run("Modify true", func(t *testing.T) {
		testValues := [][]string{[]string{"Hello World", "Test"}}
		wantValues := [][]string{[]string{"Test", "Hello World"}}

		for i, test := range testValues {
			if got := ReverseArrayString(test, true); !reflect.DeepEqual(got, wantValues[i]) {
				t.Errorf("Wanted %v but got %v", wantValues[i], got)
			}

			if !reflect.DeepEqual(testValues[i], wantValues[i]) {
				t.Errorf("Original value was not modified")
			}
		}
	})

	t.Run("Modify false", func(t *testing.T) {
		testValues := [][]string{[]string{"Hello World", "Test"}}
		wantValues := [][]string{[]string{"Test", "Hello World"}}

		for i, test := range testValues {
			if got := ReverseArrayString(test, false); !reflect.DeepEqual(got, wantValues[i]) {
				t.Errorf("Wanted %v but got %v", wantValues[i], got)
			}

			if reflect.DeepEqual(testValues[i], wantValues[i]) {
				t.Errorf("Original value was modified")
			}
		}
	})
}
