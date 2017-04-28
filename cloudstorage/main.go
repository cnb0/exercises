package main

import "fmt"

func main() {
	mcost := 12000 * 0.02
	sum := 0.0
	for i := 30; i > 0; i-- {
		sum += float64(i) / float64(30)
		fmt.Println(sum)
	}
	fmt.Println(mcost * sum)
}
