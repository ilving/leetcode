/*
283. Move Zeroes
Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Note that you must do this in-place without making a copy of the array
*/
package main

import "fmt"

func moveZeroes(nums []int) {
	if len(nums) <= 1 {
		return
	}

	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}

func main() {
	t := [][]int{
		{0, 1, 0, 3, 12},
		{0},
	}

	for _, s := range t {
		moveZeroes(s)
		fmt.Printf("%v\n", s)
	}
}
