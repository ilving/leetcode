/*
https://leetcode.com/problems/reverse-integer/

Given a signed 32-bit integer x, return x with its digits reversed.
If reversing x causes the value to go outside the signed 32-bit integer range [-2^31, 2^31 - 1], then return 0.
*/

package main

import (
	"fmt"
)

func reverse(x int) int {
	var inv_l, inv_h uint32 // 16-bits part of inverted number

	neg := 1

	if x < 0 {
		neg = -1
		x = -x
	}

	for x > 0 {
		rest := uint32(x % 10)
		x /= 10

		inv_h = inv_h<<3 + inv_h<<1        // high part * 10
		inv_l = inv_l<<3 + inv_l<<1 + rest // low * 10 + rest
		if inv_l&0xFFFF0000 > 0 {          // lower part hav more than 16 bits....
			inv_h += inv_l >> 16 // pass higest 16 bits to higer part
			inv_l &= 0xFFFF      // and clear them
		}
		if inv_h&0xFFFF8000 > 0 { // if higth part contains more than 16 bits - we cannot store inversed part
			return 0
		}
	}

	return int((inv_h&0xFFFF)<<16+(inv_l&0xFFFF)) * neg
}

func main() {
	x := []int{1563847412, -2147483412, 1534236469}

	// 9646324351 | 2147483647
	for _, v := range x {

		fmt.Printf("%d | %d\n", v, reverse(v))
	}
}
