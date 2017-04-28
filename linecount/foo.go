package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var count int
	r := bufio.NewScanner(os.Stdin)
	for r.Scan() {
		fmt.Println(r.Text())
		count++
	}
	fmt.Printf("Scanned %d lines\n", count)
}
