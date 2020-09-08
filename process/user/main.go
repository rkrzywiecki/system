package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("User ID:", os.Getuid())
	fmt.Println("Group ID:", os.Getegid())
	groups, err := os.Getgroups()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Group ID:", groups)
}
