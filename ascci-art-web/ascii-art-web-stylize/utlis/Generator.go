package docker

import (
	"strings"
)

func Generator(text, banner string) (string, error) {
	input := text

	data, err := HandleBanner(banner)
	if err != nil {
		return "", err
	}

	if err := IsValidInput(input, data); err != nil {
		return "", err
	}

	lines := strings.Count(input, "\r\n")
	words := strings.Split(input, "\r\n")

	final := ""
	for j := 0; j < len(words); j++ {
		if words[j] == "" {
			if lines > 0 {
				final += "\n"
				lines--
			}
			continue
		}
		final += PrintLine(words[j], data)
	}
	return final, nil
}
