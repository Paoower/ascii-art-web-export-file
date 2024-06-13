package asciiart

import (
	"fmt"
	"strings"
)

func GetAscii(input, style string) (string, error) {
	bannerFile, err := GetBannerFile(style)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	lines := make([]string, 0)
	words := strings.Split(input, "\n")

	for _, word := range words {
		if word == "" {
			lines = append(lines, "")
			continue
		}
		w, err := GetWord(word, bannerFile)
		if err != nil {
			return "", err
		}
		lines = append(lines, w...)
	}

	str := ""
	for _, l := range lines {
		str += l + "\n"
	}

	return str, nil
}
