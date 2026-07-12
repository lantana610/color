package main

import "strings"

func GenerateArt(str, color, substring string, banner map[rune][]string) string {
	if str == "" {
		return ""
	}

	str = strings.ReplaceAll(str, "\n", "\\n")

	word := SplitInput(str)
	var result strings.Builder
	for i, j := range word {
		if j == "" {
			if i == len(word)-1 {
				continue
			}
			result.WriteByte('\n')
			continue
		}
		input := RenderLine(j, color, substring, banner)
		for _, value := range input {
			result.WriteString(value)
			result.WriteString("\n")
		}
	}
	return result.String()
}
