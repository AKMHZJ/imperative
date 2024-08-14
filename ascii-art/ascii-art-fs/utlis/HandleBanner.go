package ascii

import (
	"errors"
	"os"
	"strings"
)

var ErrBannerNotFound = errors.New("banner file not found")

func HandleBanner(args []string) ([]string, error) {
	bannerFile := "standard"
	if len(args) == 2 {
		bannerFile = args[1]
	}
	data, err := os.ReadFile(bannerFile + ".txt")
	splitted := strings.Split(string(data), "\n")

	if err != nil {
		return []string{}, ErrBannerNotFound
	}

	return splitted, nil
}
