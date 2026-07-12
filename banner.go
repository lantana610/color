package main

import (
	"os"
	"strings"
)

func LoadBanner(str string) (map[rune][]string, error) {
	font := make(map[rune][]string)
	data, err := os.ReadFile(str)
	if err != nil {
		return nil, err
	}
	line := strings.Split(string(data), "\n")
	
	for c := ' '; c <= '~'; c++ {
		start := (int(c) - 32) * 9
		font[c] = line[start+1 : start+9]
	}
	return font, err
}
