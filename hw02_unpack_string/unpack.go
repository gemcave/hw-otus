package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var sb strings.Builder
	for i := 0; i < len(s); i++ {
		num, err := strconv.Atoi(string(s[i]))

		if err != nil && i+1 < len(s) {
			if _, err := strconv.Atoi(string(s[i+1])); err != nil {
				sb.WriteString(string(s[i]))
				continue
			}
		}

		if err == nil {
			if i == 0 {
				return "", ErrInvalidString
			}
			if _, err := strconv.Atoi(string(s[i-1])); err == nil {
				return "", ErrInvalidString
			}
		}

		if i+1 == len(s) {
			sb.WriteString(string(s[i]))
		} else if i != 0 {
			sb.WriteString(strings.Repeat(string(s[i-1]), num))
		}
	}
	return sb.String(), nil
}
