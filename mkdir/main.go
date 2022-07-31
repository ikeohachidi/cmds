package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"strings"
)

func main() {
	modeFlag := flag.Int("m", 0777, "Mode to create directory in")
	nameFlag := flag.String("n", "", "Name of folder")
	pathFlag := flag.Bool("p", false, "Will create '/' separate paths")
	verboseFlag := flag.Bool("v", false, "Be verbose")
	flag.Parse()

	if *nameFlag == "" {
		fmt.Print("Please provide a name for the folder with the -n flag")
		os.Exit(1)
	}

	// obviously the short and easy way
	// os.MkdirAll(*nameFlag, fs.FileMode(*modeFlag))

	folderNames := strings.Split(*nameFlag, "/")
	var path string
	if *pathFlag {
		for _, folderName := range folderNames {
			path += folderName + "/"
			createFolder(path, fs.FileMode(*modeFlag), *verboseFlag)
		}
		return
	}

	if len(folderNames) > 0 {
		createFolder(folderNames[0], fs.FileMode(*modeFlag), *verboseFlag)
	}
}

func createFolder(path string, perm os.FileMode, showLog bool) {
	if showLog {
		fmt.Printf("Currently creating %v folder\n", path)
	}
	err := os.Mkdir(path, perm)
	if err != nil {
		fmt.Print(err)
	}
}
