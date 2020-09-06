package main

import (
	"fmt"
	"os"
)

func main() {
	stat, err := os.Stat("file.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println(stat.Name())
	fmt.Println(stat.Size())
	fmt.Println(stat.Sys())
}
