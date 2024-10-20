package hw02unpackstring

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "a4bc2d5e", expected: "aaaabccddddde"},
		{input: "abccd", expected: "abccd"},
		{input: "", expected: ""},
		{input: "aaa0b", expected: "aab"},
		// uncomment if task with asterisk completed
		{input: `qwe\4\5`, expected: `qwe45`},
		{input: `qwe\45`, expected: `qwe44444`},
		{input: `qwe\\5`, expected: `qwe\\\\\`},
		{input: `qwe\\\3`, expected: `qwe\3`},
		{input: `ёghhб2`, expected: `ёghhбб`},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackInvalidString(t *testing.T) {
	invalidStrings := []string{"3abc", "45", "aaa10b", `d\nabc`, `qw\ne\`, `ghh\б2`}
	for _, tc := range invalidStrings {
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}

func TestIsDigit(t *testing.T) {
	for r := '1'; r <= '9'; r++ {
		require.True(t, isDigit(r))
	}
	for r := 'a'; r <= 'z'; r++ {
		require.False(t, isDigit(r))
	}
	for r := 'A'; r <= 'Z'; r++ {
		require.False(t, isDigit(r))
	}
}

func TestIsForbidden(t *testing.T) {
	require.True(t, isForbidden([]rune{'\\', 'a'}, 0))
	require.False(t, isForbidden([]rune{'\\', '1'}, 0))
	require.False(t, isForbidden([]rune{'\\', '\\'}, 0))
}

func TestIsMultiplyDigit(t *testing.T) {
	require.True(t, isMultiplyDigit([]rune{'\\', '4', '5'}, 0, 3))
	require.False(t, isMultiplyDigit([]rune{'\\', '4', 'a'}, 0, 3))
	require.False(t, isMultiplyDigit([]rune{'\\', '4', '\\'}, 0, 3))
	require.False(t, isMultiplyDigit([]rune{'\\', '4'}, 0, 2))
}

func TestIsSingleDigit(t *testing.T) {
	require.True(t, isDigitAfter([]rune{'\\', '4'}, 0))
}
