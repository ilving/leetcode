/*
https://leetcode.com/problems/search-in-rotated-sorted-array/

There is an integer array nums sorted in ascending order (with distinct values).

Prior to being passed to your function, nums is possibly rotated at an unknown pivot index k (1 <= k < nums.length)
such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed).
For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0,1,2].

Given the array nums after the possible rotation and an integer target, return the index of target if it is in nums, or -1 if it is not in nums.

You must write an algorithm with O(log n) runtime complexity
*/
package main

import "fmt"

func search(nums []int, target int) int {
	pivot := 0
	left, right := 0, len(nums)-1

	if nums[left] > nums[right] {
		for {
			mid := (left + right) >> 1
			if mid == left {
				pivot = mid + 1
				break
			}
			if nums[left] > nums[mid] {
				right = mid
			} else {
				left = mid
			}
		}
	}

	s := [][]int{nums[0:pivot], nums[pivot:]}

	for i := 0; i < len(s); i++ {
		if len(s[i]) == 0 {
			continue
		}
		if len(s[i]) == 1 && s[i][0] == target {
			return len(s[0]) * i
		}

		if s[i][0] > target || s[i][len(s[i])-1] < target {
			continue
		}

		left, right = 0, len(s[i])-1
		for {
			if s[i][left] == target {
				return len(s[0])*i + left
			}
			if s[i][right] == target {
				return len(s[0])*i + right
			}

			mid := (left + right) >> 1

			if s[i][mid] == target {
				return len(s[0])*i + mid
			}
			if mid == left {
				return -1
			}

			if s[i][mid] > target {
				right = mid
			} else {
				left = mid
			}
		}
	}

	return -1
}

func main() {
	nums := []int{1, 3}
	//nums := []int{4, 5, 6, 7, 0, 1, 2}
	//nums := []int{1, 2, 3, 4, 5, 6, 7, 0}
	target := 2

	fmt.Printf("%+v | %d | %d", nums, target, search(nums, target))
}

/*
4 5 6 7 0 1 2
4     7     2


*/
