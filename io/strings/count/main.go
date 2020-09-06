package main

import (
	"fmt"
	"strings"
)

func main() {
	count := strings.Count("test result", "s")
	fmt.Println(count)
}
