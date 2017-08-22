package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args[1:]
	if len(args) < 1 {
		panic("You need to pass in the url as a parameter")
	}

	url := args[0]

	for {
		var text string
		fmt.Printf("%s %s ", url, ">")
		fmt.Scanln(&text)
		fmt.Println(text)
	}
}
