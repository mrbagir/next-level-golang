package main

import (
	"cmp"
	"fmt"
	"strings"
)

func main() {
	userName := cmp.Or(
		getUserNameFromInput(),
		"Anonymous",
	)
	fmt.Printf("Hello, %s\n", userName)
}

func getUserNameFromInput() string {
	fmt.Print(
		"Enter your name (or press enter to use 'Anonymous'):",
	)
	var name string
	fmt.Scanln(&name)
	return strings.TrimSpace(name)
}
