package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")


func Unpack(data string) (string, error) {
	if len(data) == 0 {
		return "", nil
	}
	s := []rune(data)
	if firstIsDigit(s[0]) {
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
			times := takeTimes(s[i+1])
			res.WriteString(strings.Repeat(string(s[i]), times))
			i++
		case matchBackslash(s[i], s[i+1]):
			switch {
			case isDigit(s[i+1]) && i+2 < count && isDigit(s[i+2]):
				times := takeTimes(s[i+2])
				res.WriteString(strings.Repeat(string(s[i+1]), times))
				i += 2
			case isDigit(s[i+1]):
				res.WriteString(string(s[i+1]))
				i++
			case isBackslash(s[i+1]) && i+2 < count && isDigit(s[i+2]):
				times := takeTimes(s[i+2])
				res.WriteString(strings.Repeat(string(s[i+1]), times))
				i += 2
			case isBackslash(s[i+1]) && isBackslash(s[i+2]):
				res.WriteString(string(s[i]))
				i += countBackslash(s, i)
				res.WriteString(string(s[i]))
				fmt.Println(i)
			default:
				return "", ErrInvalidString
			}
		default:
			res.WriteString(string(s[i]))
		}
	}
	return res.String(), nil
}

func firstIsDigit(first rune) bool {
	_, err := strconv.Atoi(string(first))
	return err == nil
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
	_, err := strconv.Atoi(string(r))
	return err == nil
}

func isBackslash(r rune) bool {
	return string(r) == "\\"
}

func takeTimes(num rune) int {
	n, _ := strconv.Atoi(string(num))
	return n
}

func matchBackslash(curr rune, next rune) bool {
	return string(curr) == "\\" && (isDigit(next) || string(next) == "\\")
}

func countBackslash(data []rune, i int) int {
	n := 0
	for string(data[i+n]) == "\\" {
		n++
	}
	return n
}
