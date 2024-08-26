package docker

import (
	"errors"
)

const (
	LineHeight      = 8
	BannerFileLines = 856
)

var (
	errNotPrintable    = errors.New("400 bad request")
	errBannerBadStruct = errors.New("400 bad request")
)

func IsValidInput(text string, data []string) error {
	for _, char := range text {
		if !IsValidChar(char) {
			return errNotPrintable
		}
		if len(data) != BannerFileLines {
			return errBannerBadStruct
		}
	}
	return nil
}
