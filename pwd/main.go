package main

import (
	"fmt"
	"os"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(wd)
}
