package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("You have to specify a path")
		return
	}
	root, err := filepath.Abs(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Files in: ", root)
	var info struct {
		files int
		dirs  int
	}
	exts := map[string]int{}

	filepath.Walk(root, func(path string, f os.FileInfo, err error) error {
		if f.IsDir() {
			info.dirs++
		} else {
			info.files++
			ext := filepath.Ext(path)
			if len(strings.TrimSpace(ext)) > 0 {
				exts[ext]++
			}
		}
		fmt.Println("Dir: ", path)

		return nil
	})

	fmt.Printf("Total files: %d in %d directories\n", info.files, info.dirs)
	for ext, count := range exts {
		fmt.Printf("Extension: %s, Count: %d \n", ext, count)
	}
}
