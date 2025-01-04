package hw09structvalidator

import (
	"log"
	"strconv"
)

func validateString(val string, constrName string, constrVal string) error {
	switch constrName {
	case "len":
		return stringLenValidation(val, constrVal)
	case "regexp":
		panic("Implement me")
	case "in":
		panic("Implement me")
	}
	return ErrValidation
}

func stringLenValidation(val string, constrVal string) error {
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
