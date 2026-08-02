package service

import (
	"errors"
	"strings"

	"github.com/go-by-oksy/go-morse-converter/pkg/morse"
)

var ErrEmptyData = errors.New("empty data")

func Convert(data string) (string, error) {
	data = strings.TrimSpace(data)
	if data == "" {
		return "", ErrEmptyData
	}

	if isMorse(data) {
		return morse.ToText(data), nil
	}

	return morse.ToMorse(data), nil
}

func isMorse(data string) bool {
	hasSignal := false

	for _, char := range data {
		switch char {
		case '.', '-':
			hasSignal = true
		case ' ', '\t', '\r', '\n':
			// Separators are valid in Morse input.
		default:
			return false
		}
	}

	return hasSignal
}
