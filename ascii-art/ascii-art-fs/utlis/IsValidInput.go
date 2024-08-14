package ascii

import (
	"errors"
)

const (
	LineHeight      = 8
	BannerFileLines = 856
)

var (
	errNotPrintable    = errors.New("you've entered non-printable character")
	errBannerBadStruct = errors.New("the banner file is not well-structured")
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
