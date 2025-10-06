package main

import (
	"strings"
)

func cleanInput(text string) []string {
	textLower := strings.ToLower(text)
	textTrim := strings.TrimSpace(textLower)
	textSlice := strings.Split(textTrim, " ")

	sl := make([]string, 0, len(textSlice))
	for _, s := range textSlice {
		if s != "" {
			sl = append(sl, strings.TrimSpace(s))
		}
	}
	return sl
}
