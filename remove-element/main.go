/*
https://leetcode.com/problems/remove-element/description/

27. Remove Element
Given an integer array nums and an integer val, remove all occurrences of val in nums in-place.
The order of the elements may be changed. Then return the number of elements in nums which are not equal to val.

Consider the number of elements in nums which are not equal to val be k, to get accepted, you need to do the following things:
Change the array nums such that the first k elements of nums contain the elements which are not equal to val.
The remaining elements of nums are not important as well as the size of nums.
Return k.
*/
package main

import "fmt"

func removeElement(nums []int, val int) int {
	l := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[i], nums[l] = nums[l], nums[i]
			l++
		}
	}
	return l
}

func main() {
	a := []struct {
		arr []int
		k   int
	}{
		{arr: []int{1, 1, 2}, k: 1},
		{arr: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, k: 1},
	}

	for _, s := range a {
		l := removeElement(s.arr, s.k)
		fmt.Printf("%+v | %d | %+v\n", s, l, s.arr[0:l])
	}
}
