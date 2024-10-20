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
	s := []rune(data)
	if first := s[0]; isDigit(first) {
		return "", ErrInvalidString
	}
	var res strings.Builder
	count := len(s)
	for i := 0; i < count; i++ {
		switch {
		case lastElem(i, count):
			res.WriteString(string(s[i]))
		case matchPattern(s[i], s[i+1]):
			if isNumber(s, i, count) {
				return "", ErrInvalidString
			}
			patternStrategy(s, &i, &res)
		case matchBackslash(s[i]):
			err := backslashStrategy(s, &i, count, &res)
			if err != nil {
				return "", err
			}
		default:
			res.WriteString(string(s[i]))
		}
	}
	return res.String(), nil
}

func backslashStrategy(s []rune, ind *int, count int, res *strings.Builder) error {
	i := *ind // for comfort work
	switch {
	case isForbidden(s, i):
		return ErrInvalidString
	case isDigit(s[i+1]):
		i = backslashDigitStrategy(s, i, count, res)
	case isMultiplyBackslash(s, i, count):
		times := takeMultiplier(s[i+2])
		res.WriteString(strings.Repeat(string(s[i+1]), times))
		i += 2
	case isBackslash(s[i+1]): // isManyBackslashes(s, i):
		res.WriteString(string(s[i+1]))
		i++
	default:
		return ErrInvalidString
	}
	*ind = i // write value to outer variable
	return nil
}

func backslashDigitStrategy(s []rune, i int, count int, res *strings.Builder) int {
	if isMultiplyDigit(s, i, count) {
		times := takeMultiplier(s[i+2])
		res.WriteString(strings.Repeat(string(s[i+1]), times))
		i += 2
	} else {
		res.WriteString(string(s[i+1]))
		i++
	}
	return i
}

func isForbidden(r []rune, i int) bool {
	next := r[i+1]
	return !isDigit(next) && !isBackslash(next)
}

func isMultiplyBackslash(s []rune, i int, count int) bool {
	next := s[i+1]
	return isBackslash(next) && isThirdDigit(s, i, count)
}

func isMultiplyDigit(s []rune, i int, count int) bool {
	next := s[i+1]
	return isDigit(next) && isThirdDigit(s, i, count)
}

func isThirdDigit(s []rune, i int, count int) bool {
	return i+2 < count && isDigit(s[i+2])
}

func patternStrategy(s []rune, i *int, res *strings.Builder) {
	times := takeMultiplier(s[*i+1])
	res.WriteString(strings.Repeat(string(s[*i]), times))
	*i++
}

func lastElem(i int, count int) bool {
	return i == count-1
}

func matchPattern(curr rune, next rune) bool {
	currIsDigit := isDigit(curr)
	nextIsDigit := isDigit(next)
	return !currIsDigit && nextIsDigit && string(curr) != "\\"
}

func isNumber(s []rune, i int, count int) bool {
	if i+2 < count {
		nextIsDigit := isDigit(s[i+1])
		thirdIsDigit := isDigit(s[i+2])
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

func matchBackslash(curr rune) bool {
	return string(curr) == "\\"
}
