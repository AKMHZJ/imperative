package docker

import (
	"errors"
	"os"
	"strings"
)

var ErrBannerNotFound = errors.New("banner not found")

func HandleBanner(bannerName string) ([]string, error) {
	bannerFile := "standard"
	file := [3]string{"standard", "shadow", "thinkertoy"}
	for i := 0; i < len(file); i++ {
		if file[i] == bannerName {
			bannerFile = bannerName
			break
		}
	}
	data, err := os.ReadFile(bannerFile + ".txt")
	splitted := strings.Split(string(data), "\n")

	if err != nil {
		return []string{}, ErrBannerNotFound
	}

	return splitted, nil
}
