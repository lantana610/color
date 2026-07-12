package main

import "strings"

func RenderLine(text, color, substring string, banner map[rune][]string) []string {
	var output []string

	var ranges [][2]int
	if substring != "" {
		searchStart := 0
		for {
			idx := strings.Index(text[searchStart:], substring)
			if idx == -1 {
				break
			}
			start := searchStart + idx
			end := start + len(substring)
			ranges = append(ranges, [2]int{start, end})
			searchStart = end
		}
	}

	for i := 0; i < 8; i++ {
		var result strings.Builder
		for j, c := range text {
			render := banner[c][i]
			highlight := false
			for _, r := range ranges {
				if j >= r[0] && j < r[1] {
					highlight = true
					break
				}
			}
			if highlight {
				result.WriteString(asciicolour(color, render))
			} else {
				result.WriteString(render)
			}
		}
		output = append(output, result.String())
	}
	return output
}
