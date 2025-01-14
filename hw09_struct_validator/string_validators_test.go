package hw09structvalidator

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"
)

type testDataString struct {
	val           string
	constraintVal string
}

// in the test suppose `validate:"in:<testDataString.constraintVal>"`.
func TestStringInValidation(t *testing.T) {
	tests := []struct {
		in          testDataString
		expectedErr error
	}{
		{in: testDataString{"success", "success"}, expectedErr: nil},
		{in: testDataString{"success", "some,success,other"}, expectedErr: nil},
		{in: testDataString{"success", ",success,"}, expectedErr: nil},
		{in: testDataString{"fail", "nothing,here"}, expectedErr: ErrValidation},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			t.Parallel()

			err := func() error {
				val := tt.in.val
				constrVal := tt.in.constraintVal
				return stringInValidation(val, constrVal)
			}()
			require.ErrorIs(t, err, tt.expectedErr)
		})
	}
}

func TestStringLenValidation(t *testing.T) {
	tests := []struct {
		in          testDataString
		expectedErr error
	}{
		{in: testDataString{"success", "7"}, expectedErr: nil},
		{in: testDataString{"fail", "50"}, expectedErr: ErrValidation},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			t.Parallel()

			err := func() error {
				val := tt.in.val
				constrVal := tt.in.constraintVal
				return stringLenValidation(val, constrVal)
			}()
			require.ErrorIs(t, err, tt.expectedErr)
		})
	}
}

func TestStringRegexpValidation(t *testing.T) {
	tests := []struct {
		in          testDataString
		expectedErr error
	}{
		{in: testDataString{"234234", "\\d+"}, expectedErr: nil},
		{in: testDataString{"234234", "^\\d+$"}, expectedErr: nil},
		{in: testDataString{"234234", "[0-9]+"}, expectedErr: nil},
		{in: testDataString{"a234234", "^\\d+$"}, expectedErr: ErrValidation},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			t.Parallel()

			err := func() error {
				val := tt.in.val
				constrVal := tt.in.constraintVal
				return stringRegexpValidation(val, constrVal)
			}()
			require.ErrorIs(t, err, tt.expectedErr)
		})
	}
}

// check regexp as is.
func TestRegexp(t *testing.T) {
	tests := []struct {
		pattern  string
		val      string
		expected bool
	}{
		{pattern: `[0-9]+`, val: "1234", expected: true},
		{pattern: "\\d+", val: "234234", expected: true},
		{pattern: `\\d+`, val: "234234", expected: false},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			result, _ := regexp.MatchString(tt.pattern, tt.val)
			require.Equal(t, tt.expected, result)
		})
	}
}
