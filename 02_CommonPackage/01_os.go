package main

import (
	"fmt"
	"os"
)

func main() {

	wd, err := os.Getwd()
	if err != nil {
		return
	}

	dir, err := os.Open(wd)
	if err != nil {
		return
	}

	files, err := dir.ReadDir(-1)
	if err != nil {
		return
	}

	for _, file := range files {
		name := file.Name()
		if file.IsDir() {
			fmt.Printf("%s - 目录\n", name)
		} else {
			fmt.Printf("%s - 文件\n", name)
		}
	}
}
