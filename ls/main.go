package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
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
		if dir.IsDir() {
			printOut += fmt.Sprintf("%s ", color.CyanString(dir.Name()))
		} else {
			printOut += fmt.Sprintf("%s  ", dir.Name())
		}
	}

	fmt.Println(printOut)
}
