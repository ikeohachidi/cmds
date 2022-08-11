package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	path := flag.String("p", "", "Path to file to read") 
	lines := flag.Int("n", 10, "Number of lines to read")
	flag.Parse()

	fileInfo, err := os.Stat(*path)
	if err != nil {
		log.Fatal(err)
	}

	if fileInfo.IsDir() {
		log.Fatal("Not a folder")
	}

	file, err := os.Open(*path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileData, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	buf := bytes.NewBuffer(fileData)
	if buf.Len() < *lines {
		log.Fatal("file not long enough to read tail")
	}
	fileLines := strings.Split(buf.String(), "\n")[*lines:]

	fmt.Println(strings.Join(fileLines, "\n"))
}