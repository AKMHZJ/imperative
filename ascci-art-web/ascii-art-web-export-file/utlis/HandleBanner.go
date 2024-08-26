package docker

import (
	"errors"
	"os"
	"strings"
)

var ErrBannerNotFound = errors.New("banner not found")

func HandleBanner(bannerName string) ([]string, error) {
	file := []string{"standard", "shadow", "thinkertoy"}

	if !ArrayHas(file, bannerName) {
		return nil, ErrBannerNotFound
	}
	data, err := os.ReadFile("./handlers/" + bannerName + ".txt")
	if err != nil {
		return nil, ErrBannerNotFound
	}
	splitted := strings.Split(string(data), "\n")
	return splitted, nil
}
