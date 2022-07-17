package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

type CmdFlags struct {
	all  bool
	list bool
}

func main() {
	allFlag := flag.Bool("a", false, "Display all contents of folder including hidden items")
	listFlag := flag.Bool("l", false, "Display all information i can possibly code about the directory")
	pathFlag := flag.String("p", ".", "")
	flag.Parse()

	dirs, err := os.ReadDir(*pathFlag)
	if err != nil {
		log.Fatal(err)
	}

	flags := CmdFlags{*allFlag, *listFlag}

	var printOut strings.Builder

	for _, dir := range dirs {
		if strings.HasPrefix(dir.Name(), ".") {
			if *allFlag {
				printOut.WriteString(strBuilder(flags, dir))
			}
		} else {
			printOut.WriteString(strBuilder(flags, dir))
		}
	}

	fmt.Println(printOut.String())
}

func fileSize(dir fs.DirEntry) string {
	info, err := dir.Info()
	if err != nil {
		return "nil"
	}

	perm := info.Mode().Perm().String()
	size := strconv.Itoa(int(info.Size()/1000)) + "KB"

	var str strings.Builder
	str.WriteString(perm)
	str.WriteByte('\t')
	str.WriteString(size)

	return str.String()
}

func strBuilder(flags CmdFlags, dir fs.DirEntry) string {
	var str strings.Builder

	if flags.list {
		str.WriteString(fileSize(dir) + "\t")
	}

	if dir.IsDir() {
		str.WriteString(color.CyanString(dir.Name()) + "\t")
	} else {
		str.WriteString(dir.Name() + "\t")
	}

	if flags.list {
		str.WriteString("\n")
	}

	return str.String()
}
