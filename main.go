package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {

	colorFlag := flag.String("color", "reset", "color to use")
	flag.Parse()

	args := flag.Args()
	if len(args) < 2  {
		fmt.Println("usage: go run . --color=<color> <substring> <string>")
		return
	}

	banner := "standard.txt"
	fmt.Println(len(os.Args))
	if len(os.Args) == 5 {
		if !strings.Contains(os.Args[4], ".txt") {
			banner = os.Args[4] + ".txt"
		} else {
			banner = os.Args[4]
		}

	}
	bannerfile := banner
	substring := args[0]
	last := args[1]
	input, err := Validatetext(last)
	if err != nil {
		fmt.Printf("invalid %v is not a valid ascii", input)
	}

	txt, err := LoadBanner(bannerfile)
	if err != nil {
		fmt.Println("invalid")
		return
	}
	lasst := GenerateArt(last, *colorFlag, substring, txt)
	fmt.Println(lasst)
}
