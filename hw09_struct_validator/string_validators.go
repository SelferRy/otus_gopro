package hw09structvalidator

import (
	"log"
	"log/slog"
	"regexp"
	"strconv"
	"strings"
)

func validateString(val, constrName string, constrVal string) error {
	switch constrName {
	case "len":
		return stringLenValidation(val, constrVal)
	case "regexp":
		return stringRegexpValidation(val, constrVal)
	case "in":
		return stringInValidation(val, constrVal)
	}
	return ErrValidation
}

func stringInValidation(val, constrVal string) error {
	sep := ","
	inList := strings.Split(constrVal, sep)
	for _, target := range inList {
		if val == target {
			return nil
		} else {
			return ErrValidation
		}
	}
	slog.Info(
		`Seems something wrong. Can be 'validate:"in:"' case. Check in-statement of validation.`,
		slog.String("val", val),
		slog.String("constrVal", val),
	)
	return ErrValidation
}

func stringRegexpValidation(val, constrVal string) error {
	match, err := regexp.MatchString(val, constrVal)
	if err != nil {
		log.Fatal("regexp.MatchString(val, constrVal) was broken.")
	}
	if match {
		return nil
	}
	return ErrValidation
}

func stringLenValidation(val, constrVal string) error {
	requireLen, err := strconv.Atoi(constrVal)
	if err != nil {
		log.Fatal("strconv.Atoi(constrVal) was broken.")
	}
	valLen := len([]rune(val))
	if valLen == requireLen {
		return nil
	} else {
		return ErrValidation
	}
}
