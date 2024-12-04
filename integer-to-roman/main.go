/**/

package main

import "fmt"

func intToRoman(num int) string {
	type pos struct {
		lim int
		val string
	}
	lims := []pos{
		{lim: 1000, val: "M"},
		{lim: 900, val: "CM"},
		{lim: 500, val: "D"},
		{lim: 400, val: "CD"},
		{lim: 100, val: "C"},
		{lim: 90, val: "XC"},
		{lim: 50, val: "L"},
		{lim: 40, val: "XL"},
		{lim: 10, val: "X"},
		{lim: 9, val: "IX"},
		{lim: 5, val: "V"},
		{lim: 4, val: "IV"},
		{lim: 1, val: "I"},
	}

	s := ""
	i := 0
	for num > 0 {
		if num < lims[i].lim {
			i++
			continue
		}
		s += lims[i].val
		num -= lims[i].lim
	}

	return s
}

func main() {
	t := map[int]string{
		1:    "I",
		4:    "IV",
		3749: "MMMDCCXLIX",
		58:   "LVIII",
		1994: "MCMXCIV",
	}

	for k, v := range t {
		s := intToRoman(k)
		fmt.Printf("%5d | %-10s | %v\n", k, s, v == s)
	}
}
