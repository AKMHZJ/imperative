package ascii

import (
	"errors"
	"os"
	"strings"
)

var ErrBannerNotFound = errors.New("banner file not found")

func HandleBanner(bannerName string) ([]string, error) {
	data, err := os.ReadFile(bannerName + ".txt")
	splitted := strings.Split(string(data), "\n")

	if err != nil {
		return []string{}, ErrBannerNotFound
	}

	return splitted, nil
}
