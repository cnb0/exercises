package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: reverse2 <string>")
		return
	}

	a := []rune(os.Args[1])
	j := len(a) - 1
	for i := 0; i < j; i++ {
		a[i], a[j] = a[j], a[i]
		j--
	}
	fmt.Println(string(a))
}
