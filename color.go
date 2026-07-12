package main

 import "fmt"

func asciicolour(color, txt string) string {

	colour := map[string]string{
		"red":     "\033[31m",
		"reset":   "\033[0m",
		"green":   "\033[32m",
		"magenta": "\033[35m",
		"yellow":  "\033[33m",
		"cyan":    "\033[36m",
		"blue":    "\033[34m",
	}
	colored := colour[color]

	val := fmt.Sprintf("%s%s%s", colored, txt, "\033[0m")
	return val

}