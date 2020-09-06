package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	count, err := CopyFile("file1.txt", "file.txt")
	if err != nil {
		fmt.Println("Errr: ", err)
		return
	}
	fmt.Println(count)
}

func CopyFile(from, to string) (int64, error) {
	src, err := os.Open(from)
	if err != nil {
		return 0, err
	}
	defer src.Close()
	dst, err := os.OpenFile(to, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return 0, nil
	}
	defer dst.Close()
	return io.Copy(dst, src)
}
