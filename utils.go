package main

import (
	"github.com/atouba/piscine"
)

// Changed BasicAtoi and ConvertBase in piscine repo, but when 
// the repo was syncronized the changes were gone, which resulted
// in a bug that was mysterious to figure out the source of, now it's
// resolved (temporarily, in the event of piscine github being syncronized
// again the same bug will occur)

// index of first numeric char
func index(s string) int {
	i := 0
	for ; i < piscine.StrLen(s) ; i++ {
		if piscine.IsNumeric(string(s[i])) {
			return i
		}
	}
	return i
}

func customizedAtoi(s string) int {
	var res int = 0

	for i := index(s); i < piscine.StrLen(s); i++ {
		if !piscine.IsNumeric(string(s[i])) {
			return res
		} else {
			res += int(rune(s[i]) - '0')
			if i < piscine.StrLen(s)-1 && piscine.IsNumeric(string(s[i + 1])) {
				res *= 10
			}
		}
	}

	return res
}
