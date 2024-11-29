package main

import "fmt"

func isMonotonic(nums []int) bool {
	if len(nums) <= 2 {
		return true
	}

	diff := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			continue
		}
		if nums[i] > nums[i+1] {
			if diff > 0 {
				return false
			}
			diff = -1
		} else if nums[i] < nums[i+1] {
			if diff < 0 {
				return false
			}
			diff = 1
		}
	}

	return true
}

func main() {
	a := []int{11, 11, 9, 4, 3, 3, 3, 1, -1, -1, 3, 3, 3, 5, 5, 5}

	fmt.Printf("%+v | %v", a, isMonotonic(a))
}
