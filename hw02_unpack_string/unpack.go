package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(data string) (string, error) {
	if len(data) == 0 {
		return "", nil
	}
	r := []rune(data)
	if first := r[0]; isDigit(first) {
		return "", ErrInvalidString
	}
	var res strings.Builder
	count := len(r)
	for i := 0; i < count; i++ {
		currRune := r[i]
		switch {
		case lastElem(i, count):
			res.WriteString(string(currRune))
		case matchMultiplier(currRune, r[i+1]):
			if isNumber(r, i, count) {
				return "", ErrInvalidString
			}
			nextRune := r[i+1]
			times := takeMultiplier(nextRune)
			res.WriteString(strings.Repeat(string(currRune), times))
			i++
		case isBackslash(currRune):
			if isForbidden(r, i) {
				return "", ErrInvalidString
			}
			step, times := backslashStrategy(r, i, count)
			nextRune := r[i+1]
			res.WriteString(strings.Repeat(string(nextRune), times))
			i += step
		default:
			res.WriteString(string(currRune))
		}
	}
	return res.String(), nil
}

func backslashStrategy(r []rune, i int, count int) (int, int) {
	var step, times int
	nextRune := r[i+1]
	switch {
	case isDigit(nextRune):
		step, times = backslashDigitStrategy(r, i, count)
	case isMultiplyBackslash(r, i, count):
		times = takeMultiplier(r[i+2])
		step = 2
	case isBackslash(nextRune): // many backslashes case
		times = 1
		step = 1
	}
	return step, times
}

func isForbidden(r []rune, i int) bool {
	next := r[i+1]
	return !isDigit(next) && !isBackslash(next)
}

func backslashDigitStrategy(r []rune, i int, count int) (int, int) {
	var step, times int
	if isMultiplyDigit(r, i, count) {
		times = takeMultiplier(r[i+2])
		step = 2
	} else {
		times = 1
		step = 1
	}
	return step, times
}

func isMultiplyDigit(r []rune, i int, count int) bool {
	next := r[i+1]
	return isDigit(next) && isThirdDigit(r, i, count)
}

func isMultiplyBackslash(r []rune, i int, count int) bool {
	next := r[i+1]
	return isBackslash(next) && isThirdDigit(r, i, count)
}

func isThirdDigit(r []rune, i int, count int) bool {
	return i+2 < count && isDigit(r[i+2])
}

func lastElem(i int, count int) bool {
	return i == count-1
}

func matchMultiplier(curr rune, next rune) bool {
	return !isDigit(curr) && isDigit(next) && string(curr) != "\\"
}

func isNumber(r []rune, i int, count int) bool {
	if i+2 < count {
		nextIsDigit := isDigit(r[i+1])
		thirdIsDigit := isDigit(r[i+2])
		return nextIsDigit && thirdIsDigit
	}
	return false // length is too short
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func isBackslash(r rune) bool {
	return string(r) == "\\"
}

func takeMultiplier(num rune) int {
	n, _ := strconv.Atoi(string(num))
	return n
}
