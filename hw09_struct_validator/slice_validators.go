package hw09structvalidator

import (
	"fmt"
	"reflect"
)

func validateSlice(val reflect.Value, constrName, constrVal string) error {
	var err error
	for i := 0; i < val.Len(); i++ {
		fmt.Println(val.Type(), val.Index(i))
		v := val.Index(i)
		switch v.Kind() {
		case reflect.String:
			err = validateString(v.String(), constrName, constrVal)
		case reflect.Int:
			err = validateInt(v.Int(), constrName, constrVal)
		default:
			return ErrValidation
		}
		if err != nil {
			return ErrValidation
		}
	}
	return nil
}
