package main

import (
	"fmt"
	"strings"
)

func main() {
	str := `asdf
asdf
asdf
asdf1
asdf`
	fmt.Println(strings.Replace(str, "\n", "\\n", -1))
}
