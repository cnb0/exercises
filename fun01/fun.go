package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	for {
		s, err := r.ReadString([]byte(" ")[0])
		if err == io.EOF {
			break
		}
		if unicode.IsSpace(int(s)) {
			continue
		}
		fmt.Println(s)
	}
}
