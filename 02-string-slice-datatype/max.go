package main

import (
	"fmt"
)

func main() {
	nums := []int{16, 8, 42, 4, 23, 15}
	fmt.Println(nums)

	var max int = nums[0]
	for i := range nums {
		if nums[i] > max {
			max = nums[i]
		}
	}

	fmt.Println(max)
}
