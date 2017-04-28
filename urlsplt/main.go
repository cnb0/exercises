package main

import (
	"fmt"
	"strings"
)

func main() {
	url := "https://bots.ciscopipeline.io/vallard/rick-james"
	s := strings.Split(url, "/")
	// we want the last two strings.
	suffix := strings.Join(s[len(s)-2:], "/")
	fmt.Println(suffix)
}
