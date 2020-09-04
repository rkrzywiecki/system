package main

import (
	"fmt"
	"path/filepath"
)

func main() {

	path1, err := filepath.Abs("Test")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Dir: ", path1)

	path2 := filepath.Base(path1)
	fmt.Println("Dir: ", path2)

	path3 := filepath.Dir(path1)
	fmt.Println("Dir: ", path3)

	ext := filepath.Ext("README.md")
	fmt.Println("Ext: ", ext)

	join := filepath.Join("Test", "first")
	fmt.Println("Joined: ", join)

	list := filepath.SplitList(path1 + ":" + path3)
	for i, l := range list {
		fmt.Println(i, l)
	}

}
