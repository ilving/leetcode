/*
20. Valid Parentheses
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.
*/

package main

import "fmt"

func isValid(s string) bool {
	stack := make([]byte, 0, len(s))

	pairs := map[byte]byte{'}': '{', ')': '(', ']': '['}

	for i := 0; i < len(s); i++ {
		if last, ok := pairs[s[i]]; ok {
			if len(stack) >= 1 && stack[len(stack)-1] == last {
				stack = stack[0 : len(stack)-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, s[i])
		}
	}

	return !(len(stack) > 0)
}

func main() {
	tc := map[string]bool{
		"([)]":   false,
		"{}":     true,
		"()[]{}": true,
		"(]":     false,
	}

	for l, r := range tc {
		fmt.Printf("%-8s | %-5v | %-5v\n", l, r, isValid(l))
	}
}
