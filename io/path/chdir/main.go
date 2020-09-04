package main

import (
	"fmt"
	"os"
)

func main() {
	printCurrDir()
	changeDir()
	printCurrDir()
}

// Change current working dir
func changeDir() {
	err := os.Chdir("../..")
	if err != nil {
		fmt.Println(err)
		return
	}
}

// Print current working dir
func printCurrDir() {
	curr, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("CurrDir: ", curr)
}
