package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var sb strings.Builder
	sliceRune := []rune(s)

	for i := 0; i < len(sliceRune); i++ {
		num, err := strconv.Atoi(string(sliceRune[i]))

		if err != nil && i+1 < len(sliceRune) {
			if _, err := strconv.Atoi(string(sliceRune[i+1])); err != nil {
				sb.WriteRune(sliceRune[i])
				continue
			}
		}

		if err == nil {
			if i == 0 {
				return "", ErrInvalidString
			}
			if _, err := strconv.Atoi(string(sliceRune[i-1])); err == nil {
				return "", ErrInvalidString
			}
		}

		if i+1 == len(sliceRune) && err != nil {
			if err != nil {
				sb.WriteRune(sliceRune[i])
			}
		} else if i != 0 {
			sb.WriteString(strings.Repeat(string(sliceRune[i-1]), num))
		}
	}
	return sb.String(), nil
}
