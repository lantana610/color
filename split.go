package main

import (
	"strings"
)

func SplitInput(text string) []string {
	splitchar := strings.Split(text, "\\n")

	return splitchar
}
