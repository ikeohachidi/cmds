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
		// this implementation is underwhelming yes
		// but hey, I at least read the implementation
		// of os.Remove and os.RemoveAll before using
		// it so that's progress.
		rm(file)
	}
}

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