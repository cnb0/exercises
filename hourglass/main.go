package main

import "fmt"

func main() {
	initialized := false
	var topSum int
	matrix := [6][6]int{}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Scan(&matrix[i][j])
		}
	}

	for i := 0; i < len(matrix)-2; i++ {
		for j := 0; j < len(matrix[i])-2; j++ {
			tSum := matrix[i][j] + matrix[i][j+1] + matrix[i][j+2]
			mSum := matrix[i+1][j+1]
			bSum := matrix[i+2][j] + matrix[i+2][j+1] + matrix[i+2][j+2]
			total := tSum + mSum + bSum
			if initialized == false {
				initialized = true
				topSum = total
			}
			if total > topSum {
				topSum = total
			}
		}
	}
	fmt.Println(topSum)
}
