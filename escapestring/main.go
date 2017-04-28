package main

import (
	"fmt"
	"html"
)

func main() {
	str := "this is a string \" with \\ escaped quotes \" \""
	s := html.EscapeString(str)
	fmt.Println(html.UnescapeString(s))
}
