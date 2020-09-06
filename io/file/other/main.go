package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	stat, _ := os.Stat("file.txt")
	fmt.Println(stat.Mode())

	if err := os.Chmod("file.txt", 06444); err != nil {
		fmt.Println(err)
		return
	}
	stat, _ = os.Stat("file.txt")
	fmt.Println(stat.Mode())

	mtime := time.Date(2011, 10, 1, 12, 30, 0, 0, time.UTC)
	atime := time.Date(2011, 8, 1, 12, 30, 0, 0, time.UTC)

	err := os.Chtimes("file.txt", atime, mtime)
	if err != nil {
		fmt.Println(err)
	}
}
