package main

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
)

func main() {
	uid := os.Getuid()
	u, err := user.LookupId(strconv.Itoa(uid))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("User %s (uid %d)\n", u.Username, uid)
	gid := os.Getegid()
	group, err := user.LookupGroupId(strconv.Itoa(gid))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Group: %s (uid %d)\n", group.Name, gid)
}
