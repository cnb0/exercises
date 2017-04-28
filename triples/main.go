package main

import "fmt"

func main() {
	var i [6]int
	for c := 0; c < len(i); c++ {
		fmt.Scan(&i[c])
	}
	var aCount, bCount int
	for triple := 0; triple < 3; triple++ {
		if i[triple] > i[triple+3] {
			aCount++
		} else if i[triple] < i[triple+3] {
			bCount++
		}
	}
	fmt.Println(aCount, bCount)
}
