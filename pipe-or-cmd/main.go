package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func scanLines(in io.Reader) {
	s := bufio.NewScanner(in)
	for s.Scan() {
		fmt.Println(s.Text())
	}
}

func main() {
	if len(os.Args) > 1 {
		fmt.Println("arg 1: ", os.Args[1])
		f, err := os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		scanLines(f)
	} else {
		fmt.Println("No args, reading from stdin")
		scanLines(os.Stdin)
	}
}
