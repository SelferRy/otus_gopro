package hw09structvalidator

import (
	"log"
	"log/slog"
	"regexp"
	"strconv"
	"strings"
)

func validateString(val, constrName, constrVal string) error {
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
		}
	}
	return ErrValidation
}

func stringRegexpValidation(val, pattern string) error {
	match, err := regexp.MatchString(pattern, val)
	if err != nil {
		log.Fatal("regexp.MatchString(val, pattern) was broken.")
	}
	slog.Debug(
		"regexp",
		slog.String("val", val),
		slog.String("pattern", pattern),
		slog.Any("match", match),
	)
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
