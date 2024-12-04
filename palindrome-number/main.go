/*
9. Palindrome Number

Given an integer x, return true if x is a palindrome, and false otherwise.
*/
package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	if x%10 == 0 {
		return false
	}

	for acc := 0; x > 0 && x > acc; {
		acc = acc*10 + x%10
		x /= 10

		if x == acc || x/10 == acc {
			return true
		}
	}

	return false
}

func main() {
	a := 11011

	fmt.Printf("%d | %v", a, isPalindrome(a))
}
