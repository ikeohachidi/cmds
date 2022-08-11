package main

import (
	"os"
)

func main() {
	// won't use flogs here since we'll at most 
	// always have one or more arguments all of 
	// which should be deleted
	args := os.Args[1:]

	for _, file := range args {
		// we're not going to use os.RemoveAll
		// although it also recursively deletes 
		// content too
		rm(file)
	}
}

// rm will recursively delete all contents
// of a folder before deleting the folder itself
// os.Remove will not delete a folder with content
func rm(path string) error {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return err
	}

	if !fileInfo.IsDir() {
		if err = os.Remove(path); err != nil {
			return err
		}
	} else {
		dirFiles, err := os.ReadDir(path)
		if err != nil {
			return err
		}

		for _, file := range dirFiles {
			s := path+"/"+file.Name()
			rm(s)
		}
		if err = os.Remove(path); err != nil {
			return err
		}
	}

	return nil
}