package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var i uint32 = 4
	var d float32 = 4.0
	var s string = "hackerrank "

	scanner := bufio.NewReader(os.Stdin)

	var ii uint32
	var dd float32
	var ss string

	fmt.Scanf("%d\n", &ii)
	fmt.Scanf("%g\n", &dd)
	ss, _ = scanner.ReadString('\n')
	fmt.Println(ii + i)
	fmt.Printf("%0.1f\n", d+dd)
	fmt.Print(s + ss)
}
