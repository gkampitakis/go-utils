package goUtils

import (
	"reflect"
	"testing"
)

func TestReverseArrayString(t *testing.T) {

	testValues := [][]string{[]string{"Hello World", "Test", "Blue"}, []string{"1", "2", "3"}}
	wantValues := [][]string{[]string{"Blue", "Test", "Hello World"}, []string{"3", "2", "1"}}

	for i, test := range testValues {

		if got := ReverseArrayString(test); !reflect.DeepEqual(got, wantValues[i]) {

			t.Errorf("Wanted %v but got %v", wantValues[i], got)

		}

	}

}
