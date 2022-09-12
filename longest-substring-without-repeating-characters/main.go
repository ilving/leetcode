package main

func lengthOfLongestSubstring(s string) int {
	m := [128 - 32]bool{}

	sLen := 0
	cLen := 0
	for i := 0; i < len(s); i++ {
		sym := s[i]

		if m[sym-32] {
			if cLen > sLen {
				sLen = cLen
			}
			i -= cLen
			for k := 0; k < 128-32; k++ {
				m[k] = false
			}
			cLen = 0
		} else {
			m[sym-32] = true
			cLen++
		}
	}

	if cLen > sLen {
		return cLen
	}

	return sLen
}

func main() {
	s := "dvdffdva"
	println(lengthOfLongestSubstring(s))
}
