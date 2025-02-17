package hw09structvalidator

import (
	"errors"
	"fmt"
	"log"
	"log/slog"
	"reflect"
	"strings"
)

var (
	ErrValidation = errors.New("validation error")
	ErrTagParsing = errors.New("tag parsing error")
)

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	return ErrValidation.Error()
}

func Validate(v interface{}) error {
	vType := reflect.TypeOf(v)
	vVal := reflect.ValueOf(v)
	numField := vType.NumField()
	errValid := make(ValidationErrors, 0)
	unvalidated := func(typeField reflect.StructField) bool {
		constr, _ := defineConstraints(typeField)
		return constr == nil
	}
	for i := 0; i < numField; i++ {
		typeField := vType.Field(i)
		valField := vVal.Field(i)
		if ok := unvalidated(typeField); ok {
			continue
		}
		validated, err := validateField(typeField, valField)
		if err != nil {
			log.Fatal("validation problem\n%w\n", err)
		}
		errValid = append(errValid, validated) // TODO: then add check to result
	}
	return errValid
}

func validateField(typeField reflect.StructField, valField reflect.Value) (ValidationError, error) {
	constraintMap, err := defineConstraints(typeField)
	if err != nil {
		return ValidationError{}, fmt.Errorf("defineConstraints(typeField) was broken: %w", err)
	}
	var errValid error
	for cName, cVal := range constraintMap {
		if errValid != nil {
			break
		}
		switch typeField.Type.Kind() {
		case reflect.String:
			errValid = validateString(valField.String(), cName, cVal)
		case reflect.Int:
			errValid = validateInt(valField.Int(), cName, cVal)
		case reflect.Slice:
			errValid = validateSlice(valField, cName, cVal)
		default:
			errValid = ErrValidation
		}
	}
	//nolint:nilerr
	return ValidationError{Field: typeField.Name, Err: errValid}, nil // here error always nil, no need return err
}

// make map like {"len": 5, "...": "..."}.
func defineConstraints(field reflect.StructField) (map[string]string, error) {
	tag := field.Tag.Get("validate")
	tagSep := "|"
	constraints := strings.Split(tag, tagSep)
	constraintMap := make(map[string]string, len(constraints))
	tagValSep := ":"
	for _, constr := range constraints {
		parsedTag := strings.SplitN(constr, tagValSep, 2)
		if len(parsedTag) != 2 {
			slog.Error("problem with parsedTag", slog.Any("parsedTag", parsedTag))
			return nil, ErrTagParsing
		}
		name, val := parsedTag[0], parsedTag[1]
		if v, ok := constraintMap[name]; ok {
			slog.Info(
				"validation info has duplicates. Handle only latest.",
				slog.String("latest name", name),
				slog.String("latest val", val),
				slog.String("was skipped", v),
			)
		}
		constraintMap[name] = val
	}
	return constraintMap, nil
}
