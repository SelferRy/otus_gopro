package hw09structvalidator

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

type UserRole string

// Test the function on different structures and other types.
type (
	User struct {
		ID     string `json:"id" validate:"len:36"`
		Name   string
		Age    int             `validate:"min:18|max:50"`
		Email  string          `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole        `validate:"in:admin,stuff"`
		Phones []string        `validate:"len:11"`
		meta   json.RawMessage //nolint:unused
	}

	App struct {
		Version string `validate:"len:5"`
	}

	Token struct {
		Header    []byte
		Payload   []byte
		Signature []byte
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}

	Pipe struct {
		Digits string `validate:"regexp:\\d+|len:20"`
		Range  int    `validate:"min:0|max:100"`
	}

	Period struct {
		Start string `validate:"in:2025-01-01,2025-01-02"`
		End   string `validate:"in:2025-01-01,2025-01-02"`
	}
)

func TestValidate(t *testing.T) {
	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{
			in: User{
				ID:     "id",
				Name:   "name",
				Age:    24,
				Email:  "email@gmail.com",
				Role:   "somebody",
				Phones: []string{"122113", "12312312345"},
			},
			expectedErr: ValidationErrors{
				ValidationError{"ID", ErrValidation},
				ValidationError{"Age", nil},
				ValidationError{"Email", nil},
				ValidationError{"Role", ErrValidation},
			},
		}, {
			in: App{
				Version: "1",
			},
			expectedErr: ValidationErrors{
				ValidationError{"Version", ErrValidation},
			},
		}, {
			in: Token{
				Header:    []byte{123},
				Payload:   []byte{12},
				Signature: []byte{58},
			},
			expectedErr: ValidationErrors{},
		}, {
			in: Response{
				Code: 1,
				Body: "body",
			},
			expectedErr: ValidationErrors{
				ValidationError{"Code", ErrValidation},
			},
		}, {
			in: Pipe{
				Digits: "2304234234",
				Range:  500,
			},
			expectedErr: ValidationErrors{
				ValidationError{"Digits", ErrValidation},
				ValidationError{"Range", ErrValidation},
			},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()

			err := Validate(tt.in)
			require.ErrorIs(t, err, tt.expectedErr)
			_ = tt
		})
	}
}

func TestValidateTDD(t *testing.T) {
	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{
			in: App{
				Version: "1", // `validate:"len:5"`
			},
			expectedErr: ValidationErrors{
				ValidationError{"Version", ErrValidation},
			},
		}, {
			in: App{
				Version: "11111", // `validate:"len:5"`
			},
			expectedErr: ValidationErrors{
				ValidationError{"Version", nil},
			},
		}, {
			in: Period{
				Start: "2024-12-12", // `validate:"in:2025-01-01,2025-01-02"`
				End:   "2025-01-01",
			},
			expectedErr: ValidationErrors{
				ValidationError{"Start", ErrValidation},
				ValidationError{"End", nil},
			},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			t.Parallel()

			err := Validate(tt.in)
			require.Equal(t, tt.expectedErr, err)
		})
	}
}

func TestReadDir(t *testing.T) {
	t.Run("read directory", func(t *testing.T) {
		target := ErrValidation
		require.ErrorIs(t, target, ErrValidation)
	})
}
