package main

import "fmt"

func Validatetext(s string) (rune, error) {
	for _, i := range s {
		if i < 32 || i > 126 {
			return i, fmt.Errorf("error %q", i)
		}
	}
	return 0, nil
}
