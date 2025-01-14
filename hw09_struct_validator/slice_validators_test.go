package hw09structvalidator

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

type structValidationData struct {
	data           reflect.Value
	constraintName string
	constraintVal  string
}

func refVal[T []string | []int](v T) reflect.Value {
	return reflect.ValueOf(v)
}

func TestValidateSliceLen(t *testing.T) {
	tests := []struct {
		in   structValidationData
		want error
	}{
		{structValidationData{refVal([]string{"some", "othe"}), "len", "4"},
			nil},
		{structValidationData{refVal([]string{"some", "other"}), "len", "4"},
			ErrValidation},
		{structValidationData{refVal([]int{4, 16}), "in", "4,16"},
			nil},
		{structValidationData{refVal([]int{4, 17}), "in", "4,16"},
			ErrValidation},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			t.Parallel()

			res := validateSlice(tt.in.data, tt.in.constraintName, tt.in.constraintVal)
			if !errors.Is(res, tt.want) {
				t.Fatalf(`validateSlice broken`)
			}
		})
	}
}
