/*
https://leetcode.com/problems/reverse-integer/

Given a signed 32-bit integer x, return x with its digits reversed.
If reversing x causes the value to go outside the signed 32-bit integer range [-2^31, 2^31 - 1], then return 0.
*/

package main

import "fmt"

func reverse(x int) int {
	var inv int64 = 0

	neg := 1

	if x < 0 {
		neg = -1
		x = -x
	}

	for x > 0 {
		rest := x % 10
		x = (x - rest) / 10
		inv = inv*10 + int64(rest)
	}

	if inv > (1<<31)-1 {
		return 0
	}

	return int(inv) * neg
}

func main() {
	x := 1534236469
	fmt.Printf("%d | %d", x, reverse(x))
}
