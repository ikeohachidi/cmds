package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
)

func main() {
	printOut := ""

	allFlag := flag.Bool("a", false, "Display all contents of folder including hidden items")
	pathFlag := flag.String("p", ".", "")
	flag.Parse()

	dirs, err := os.ReadDir(*pathFlag)
	if err != nil {
		log.Fatal(err)
	}

	for _, dir := range dirs {
		if strings.HasPrefix(dir.Name(), ".") {
			if *allFlag {
				printer(&printOut, dir)
			}
		} else {
			printer(&printOut, dir)
		}
	}

	fmt.Println(printOut)
}

func printer(to *string, dir fs.DirEntry) {
	if dir.IsDir() {
		*to += fmt.Sprintf("%s ", color.CyanString(dir.Name()))
	} else {
		*to += fmt.Sprintf("%s  ", dir.Name())
	}
}
