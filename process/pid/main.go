package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Current PID:", os.Getpid())
	fmt.Println("Current parent PID:", os.Getppid())
}
