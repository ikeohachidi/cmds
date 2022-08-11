package main

import (
	"flag"
	"io"
	"log"
	"os"
)


func main() {
	source := flag.String("p", "", "path to source directory")
	dest := flag.String("d", "", "path to destination directory")
	flag.Parse()

	fileInfo, err := os.Stat(*source)
	if err != nil {
		log.Fatal(err)
	}

	if fileInfo.IsDir() {
		if err := dirCopy(*source, *dest); err != nil {
			log.Fatal(err)
		}
	}

	fileCopy(*source, *dest)
}

func dirCopy(source string, dest string) error {
	if err := os.Mkdir(dest, 0777); err != nil {
		return err
	}

	dirEntry, err := os.ReadDir(source)
	if err != nil {
		return err
	}

	for _, file := range dirEntry {
		s := source+"/"+file.Name()
		d := dest+"/"+file.Name()

		if (file.IsDir()) {
			if err := dirCopy(s, d); err != nil {
				return err
			}
		} else {
			fileCopy(s, d)
		}
	}

	return nil
}

func fileCopy(source string, dest string) error {
	sourceFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()

	if _, err = io.Copy(destFile, sourceFile); err != nil {
		return err
	}

	return nil
}