// I do not use "testify/require" in the test module by design for experiments with native test library.
package hw09structvalidator

import (
	"errors"
	"fmt"
	"testing"
)

type intInput struct {
	data          int64
	constraintVal string
}

func TestIntMinValidation(t *testing.T) {
	// simple manual test with only native "testing".
	func() {
		data := int64(5)
		threshold := "4"
		result := intMinValidation(data, threshold)
		var want error
		if !errors.Is(result, want) {
			t.Fatalf(`intMinValidation(%d, %s) = %v, want %v`, data, threshold, result, want)
		}
	}()

	// table-driven tests with only native "testing".
	tests := []struct {
		in   intInput
		want error
	}{
		{intInput{15, "10"}, nil},
		{intInput{10, "15"}, ErrValidation},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			res := intMinValidation(tt.in.data, tt.in.constraintVal)
			if !errors.Is(res, tt.want) {
				t.Fatalf(`intMinValidation(%d, %s) = %v, want %v`, tt.in.data, tt.in.constraintVal, res, tt.want)
			}
		})
	}
}

func TestIntMaxValidation(t *testing.T) {
	tests := []struct {
		in   intInput
		want error
	}{
		{intInput{15, "10"}, ErrValidation},
		{intInput{10, "15"}, nil},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			res := intMaxValidation(tt.in.data, tt.in.constraintVal)
			if !errors.Is(res, tt.want) {
				t.Fatalf(`intMaxValidation(%d, %s) = %v, want %v`, tt.in.data, tt.in.constraintVal, res, tt.want)
			}
		})
	}
}
