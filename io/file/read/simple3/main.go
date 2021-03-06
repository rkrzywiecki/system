package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please specify a path")
		return
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer f.Close()
	r := bufio.NewReader(f)
	var rowCount int
	for err == nil {
		var b []byte
		for moar := true; err == nil && moar; {
			b, moar, err = r.ReadLine()
			if err == nil {
				fmt.Print(string(b))
			}
		}
		if err == nil {
			fmt.Println()
			rowCount++
		}
	}
	if err != nil && err != io.EOF {
		fmt.Println("\nError: ", rowCount)
		return
	}
	fmt.Println("\nRow count: ", rowCount)
}
