package main

import (
	"fmt"
	"os"
)

func main() {
	if err := os.Rename("file.txt", "file1.txt"); err != nil {
		fmt.Println("Error: ", err)
	}
}
