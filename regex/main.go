package main

import (
	"fmt"
	"regexp"
)

func main() {
	var validID = regexp.MustCompile(`.*-.*-.*`)
	fmt.Println(validID.MatchString("vallard - foo - blah"))
	fmt.Println(validID.MatchString("vallard2 - foo blah"))
	fmt.Println(validID.MatchString("vallard3 foo blah"))
}
