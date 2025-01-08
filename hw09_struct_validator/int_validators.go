package hw09structvalidator

import (
	"log"
	"strconv"
	"strings"
)

func validateInt(val int64, constrName, constrVal string) error {
	switch constrName {
	case "min":
		return intMinValidation(val, constrVal)
	case "max":
		return intMaxValidation(val, constrVal)
	case "in":
		return intInValidation(val, constrVal)
	}
	return ErrValidation
}

func intMinValidation(val int64, constrVal string) error {
	threshold := extractInt(constrVal)
	if val <= threshold {
		return ErrValidation
	}
	return nil
}

func intMaxValidation(val int64, constrVal string) error {
	threshold := extractInt(constrVal)
	if val >= threshold {
		return ErrValidation
	}
	return nil
}

func intInValidation(val int64, constrVal string) error {
	valMap := func() map[int64]struct{} {
		sep := ","
		valSlice := strings.Split(constrVal, sep)
		res := make(map[int64]struct{}, len(valSlice))
		for _, v := range valSlice {
			res[extractInt(v)] = struct{}{}
		}
		return res
	}()
	if _, ok := valMap[val]; !ok {
		return ErrValidation
	}
	return nil
}

// convert constrVal to int64
func extractInt(val string) int64 {
	cVal, err := strconv.Atoi(val)
	if err != nil {
		log.Fatal("constrVal does not contain int")
	}
	return int64(cVal)
}
