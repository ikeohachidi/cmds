package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
)

func main() {
	modeFlag := flag.Int("m", int(777), "Mode to create directory in")
	nameFlag := flag.String("n", "", "Name of folder")
	flag.Parse()

	if *nameFlag == "" {
		fmt.Print("Please provide a name for the folder with the -n flag")
		os.Exit(1)
	}

	os.Mkdir(*nameFlag, fs.FileMode(*modeFlag))
}
