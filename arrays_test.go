package arrays

import (
	"reflect"
	"testing"
)

func TestReverseArrayString(t *testing.T) {

	testValues := [][]string{[]string{"Hello World", "Test", "Blue"}, []string{"1", "2", "3"}}
	wantValues := [][]string{[]string{"Blue", "Test", "Hello World"}, []string{"3", "2", "1"}}

	for i, test := range testValues {

		if got := reverseArrayString(test); !reflect.DeepEqual(got, wantValues[i]) {

			t.Errorf("Wanted %v but got %v", wantValues[i], got)

		}

	}

}

func TestHelloWorld(t *testing.T) {

	want := "Hello, world."

	if got := HelloWorld(); got != want {

		t.Errorf("Wanted %v but got %v", want, got)

	}

}

func TestProverb(t *testing.T) {

	want := "Concurrency is not parallelism."

	if got := Proverb(); got != want {

		t.Errorf("Wanted %v but got %v", want, got)

	}

}
