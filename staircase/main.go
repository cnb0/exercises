package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	for i := n; i > 0; i-- {
		for j := 0; j < i-1; j++ {
			fmt.Printf(" ")
		}
		for k := 0; k < n-i+1; k++ {
			fmt.Printf("#")
		}
		fmt.Println()
	}
}
