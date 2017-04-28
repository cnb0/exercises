package main

import "fmt"

func main() {
	var size int
	fmt.Scan(&size)
	m := make([][]int, size)
	var pd, sd int
	for i := 0; i < size; i++ {
		m[i] = make([]int, size)
		for j := 0; j < size; j++ {
			fmt.Scan(&m[i][j])
			if i+j == size-1 {
				sd += m[i][j]
			}
			if i == j {
				pd += m[i][j]
			}
		}
	}
	abs := pd - sd
	if abs < 0 {
		abs = -abs
	}
	fmt.Println(abs)
}
