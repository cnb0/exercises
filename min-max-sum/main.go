package main

import "fmt"

func sortNum(nums []int64, num int64) []int64 {
	if len(nums) == 0 {
		return append(nums, num)
	}
	nums = append(nums, num)
	for j := len(nums) - 1; j > 0; j-- {
		if nums[j] < nums[j-1] {
			nums[j], nums[j-1] = nums[j-1], nums[j]
		}
	}
	return nums
}

func main() {
	var nums []int64
	var num, min, max int64
	for i := 0; i < 5; i++ {
		fmt.Scan(&num)
		nums = sortNum(nums, num)
	}
	// get max:
	max = nums[4]
	min = nums[0]
	for i := 1; i < 4; i++ {
		max += nums[i]
		min += nums[i]
	}
	fmt.Println(min, max)

}
