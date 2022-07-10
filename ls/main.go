package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	path := "."
	args := os.Args
	printOut := ""

	if len(args) > 1 {
		path = args[1]
	}

	dirs, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, dir := range dirs {
		printOut += fmt.Sprintf("%v   ", dir.Name())
	}

	log.Print(printOut)
}
