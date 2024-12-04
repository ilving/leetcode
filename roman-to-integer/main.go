/*
https://leetcode.com/problems/roman-to-integer/description/

13. Roman to Integer
Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
For example, 2 is written as II in Roman numeral, just two ones added together.
12 is written as XII, which is simply X + II.
The number 27 is written as XXVII, which is XX + V + II.

Roman numerals are usually written largest to smallest from left to right.
However, the numeral for four is not IIII. Instead, the number four is written as IV.
Because the one is before the five we subtract it making four.
The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

I can be placed before V (5) and X (10) to make 4 and 9.
X can be placed before L (50) and C (100) to make 40 and 90.
C can be placed before D (500) and M (1000) to make 400 and 900.
Given a roman numeral, convert it to an integer.
*/
package main

import "fmt"

func romanToInt(s string) int {
	vals := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

	if s == "" {
		return 0
	}

	ints := make([]int, len(s))
	for i, sym := range []byte(s) {
		ints[i] = vals[sym]
	}

	res := ints[0]
	for i := 1; i < len(ints); i++ {
		res += ints[i]
		if ints[i-1] < ints[i] {
			res = res - 2*ints[i-1]
		}
	}

	return res
}

func main() {
	s := []string{
		"III",     // 3
		"IV",      // 4
		"V",       // 5
		"VI",      // 6
		"LVIII",   // 58
		"MCMXCIV", // 1994 M CM XC IV
	}

	for _, v := range s {
		fmt.Printf("%-10s | %4d\n", v, romanToInt(v))
	}
}
