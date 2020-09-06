package main

import (
	"fmt"
	"os"
)

func main() {
	if err := os.Truncate("file.txt", 4096); err != nil {
		fmt.Println("Error: ", err)
	}
}
