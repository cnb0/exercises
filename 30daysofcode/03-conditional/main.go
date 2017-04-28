package main

import "fmt"

func main() {
	var i int
	fmt.Scanf("%d", &i)
	if (i % 2) == 1 {
		fmt.Println("Weird")
		return
	}

	if i < 6 {
		fmt.Println("Not Weird")
		return
	}

	if i < 21 {
		fmt.Println("Weird")
	} else {
		fmt.Println("Not Weird")
	}

}
