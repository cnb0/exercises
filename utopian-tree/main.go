package main

import "fmt"

func main() {
	// number of test cases
	var t int
	fmt.Scanf("%d", &t)
	n := make([]int, t)
	for i := 0; i < t; i++ {
		fmt.Scanf("%d", &n[i])
	}
	// now go through test cycles.
	for i := 0; i < len(n); i++ {
		fmt.Printf("%d\n", treegrow(n[i]))
	}
}

// use recursion
func treegrow(n int) int {
	if n == 0 {
		return 1
	}
	if (n % 2) == 0 {
		return 1 + treegrow(n-1)
	}
	return 2 * treegrow(n-1)
}
