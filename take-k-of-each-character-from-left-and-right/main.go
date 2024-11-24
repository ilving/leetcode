/*
https://leetcode.com/problems/take-k-of-each-character-from-left-and-right/description/?envType=daily-question&envId=2024-11-20.

You are given a string s consisting of the characters 'a', 'b', and 'c' and a non-negative integer k.
Each minute, you may take either the leftmost character of s, or the rightmost character of s.

Return the minimum number of minutes needed for you to take at least k of each character, or return -1 if it is not possible to take k of each character.
*/

package main

import (
	"fmt"
)

type calc []int

func (c calc) inCondition(k int) bool {
	r := 0
	for _, v := range c {
		if v >= k {
			r++
		}
	}
	return r == len(c)
}

func takeCharacters(s string, k int) int {
	letters := calc{0, 0, 0}
	for i := 0; i < len(s); i++ {
		letters[s[i]-'a']++
	}

	if !letters.inCondition(k) {
		return -1
	}

	steps := len(s)
	winLen := 1

	if steps == len(letters)*k {
		return steps
	}

	left, right := 0, 0

	for ; right < len(s); right++ {
		letters[s[right]-'a']--

		for ; left < right && !(letters[0] >= k && letters[1] >= k && letters[2] >= k); left++ {
			letters[s[left]-'a']++
		}

		if letters[0] >= k && letters[1] >= k && letters[2] >= k {
			winLen = max(winLen, right-left+1)
		}
	}

	return steps - winLen
}

func main() {
	s := "aabcaaacaabc"
	k := 2

	fmt.Printf("%s | %d | %d", s, k, takeCharacters(s, k))
}

/*
0 1 2 3 4 5 6
a b c a b c c - 2,2,3,7
[0:0]1,2,3-
[0:1]1,1,3-
[1:1]2,1,3-
[1:2]2,1,2-
[2:2]2,2,2,(1/6)+
[2:3]1,2,2,-
[3:3]1,2,3,-
[3:4]1,1,3,-
[4:4]2,1,3,-
[4:5]2,1,2,-
[5:5]2,2,2,(1/6)=
[5:6]2,2,1,-
[6:6]2,2,2,(1/6)=
*/

/*
0 1 2 3 4 5 6 7 8 9 A B
a a b a a a a c a a b c - 8,2,2,12

[0:0]7,2,2,11+
[0:1]6,2,2,10+
[0:2]6,1,2,-
[1:2]7,1,2,-
[2:2]8,1,2,-
[2:3]7,1,2,-
[3:3]7,2,2,11>
[3:4]6,2,2,10=
[3:5]5,2,2,9+
[3:6]4,2,2,8+
[3:7]4,2,1,-
[4:7]5,2,1,-
[5:7]6,2,1,-
[6:7]7,2,1,-
[7:7]8,2,1,-
[7:8]7,2,1,-
[8:8]7,2,2,11>
[8:9]6,2,2,10>
[8:A]6,1,2,-
[9:A]7,1,2,-
[A:A]8,1,2,-
[A:B]8,1,1,-
[B:B]8,2,1,-
*/
/*
0 1 2 3 4 5 6 7 8 9 A B
a a b c a a a c a a b c - 7,2,3,12

[0:0]6,2,3,11+
[0:1]5,2,3,10+
[0:2]5,1,3,-
[1:2]6,1,3,-
[2:2]7,1,3,-
[2:3]7,1,2,-
[3:3]7,2,2,11>
[3:4]6,2,2,10>
[3:5]5,2,2,9+
[3:6]4,2,2,8+
[3:7]4,2,1,-
[4:7]4,2,2,8>
[4:8]3,2,2,7+
[4:9]2,2,2,6+
[4:A]2,1,2,-
[5:A]3,1,2,-
[6:A]4,1,2,-
[7:A]4,1,3,-
[8:A]5,1,3,-
[9:A]6,1,3,-
[A:A]7,1,3,-
[A:B]7,1,2,-
[B:B]7,2,2,11>

*/
