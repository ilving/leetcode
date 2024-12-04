/*
26. Remove Duplicates from Sorted Array
Given an integer array nums sorted in non-decreasing order,
remove the duplicates in-place such that each unique element appears only once.
The relative order of the elements should be kept the same. Then return the number of unique elements in nums.

Consider the number of unique elements of nums to be k, to get accepted, you need to do the following things:

Change the array nums such that the first k elements of nums contain the unique elements in the order they were present in nums initially. The remaining elements of nums are not important as well as the size of nums.
Return k.
*/

package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	l := 0
	p := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] != p {
			l++
			p = nums[i]
			nums[l] = p
			continue
		}

	}

	return l + 1
}

func main() {
	a := [][]int{
		{1, 1, 2},
		{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
	}

	for _, s := range a {
		l := removeDuplicates(s)
		fmt.Printf("%+v | %d | %+v\n", s, l, s[0:l])
	}
}
