package main

import (
	"fmt"
	"os"
)

func main() {
	word := os.Args[1]

	switch l := len(word); {
	case word == "hi":
		fmt.Println("Very informal!")
		fallthrough
	case word == "hello":
		fmt.Println("Hi!")
	case l == 1:
		fmt.Println("One letter word?")
	case 1 < l && l < 10, word == "ok":
		fmt.Println("ok or 2-9 chars word")
	case word == "goodbye",  word == "bye":
		fmt.Println("See you soon!")
	default:
		fmt.Println("What?")
	}
}
