package main

//see: https://visualgo.net/sorting for visualization
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var arr []int
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		i, err := strconv.Atoi(s.Text())
		if err != nil {
			fmt.Println(err)
		}
		arr = append(arr, i)
	}
	quicksort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

// 20 9 923 38 398
func partition(arr []int, lo int, hi int) int {
	pivot := arr[hi]
	i := lo
	// go through each item and move everything less than pivot down.
	for j := lo; j < hi; j++ {
		if arr[j] <= pivot {
			arr[j], arr[i] = arr[i], arr[j]
			i++
		}
	}
	arr[hi], arr[i] = arr[i], arr[hi]
	return i
}

// quicksort modifies the array in place
func quicksort(arr []int, lo int, hi int) {
	if lo < hi {
		p := partition(arr, lo, hi)
		quicksort(arr, lo, p-1)
		quicksort(arr, p+1, hi)
	}
	fmt.Println(arr)
}
