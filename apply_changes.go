package main

import (
	"github.com/atouba/piscine"
)

func isVowelStart(word string) bool {
	vowels := []rune{'a', 'e', 'i', 'o', 'h'}
	for _, v := range vowels {
		if rune(word[0]) == v || rune(word[0]) == v - 32 {
			return true
		}
	}
	return false
}

func applyChanges(tokens []string) {
	for i, token := range tokens {
		if token[0] == '(' {
			applyCaseNumericConversion(tokens, i)
		}
		if ((token == "a" || token == "A") && i + 1 < len(tokens)) &&
		((isVowelStart(tokens[i + 1])) || 
		(tokens[i + 1][0] ==  '(' &&
		i + 2 < len(tokens) && isVowelStart(tokens[i + 2]))) {
			tokens[i] = tokens[i] + "n"
		}
	}
}

func rightFunction(c rune) (func (string) (string)) {
	convert_case_funcs := []func(string) (string) {piscine.ToUpper, piscine.ToLower, piscine.Capitalize}
	bases := []string {"01", "0123456789ABCDEF", "0123456789"}
	var f func (string) (string)
	switch c {
	case 'l':
		f = convert_case_funcs[1]
	case 'u' :
		f = convert_case_funcs[0]
	case 'c' :
		f = convert_case_funcs[2]
	case 'b' :
		f = func (n string) string {
			return piscine.ConvertBase(n, bases[0], bases[2])
		}
	default:
		f = func (n string) string {
			return piscine.ConvertBase(piscine.ToUpper(n), bases[1], bases[2])
		}
	}
	return f
}

func applyCaseNumericConversion(tokens []string, index int) {
	var exec func (string) (string)
	var number_of_words int

	number_of_words = max(1, customizedAtoi(tokens[index]))
	
  exec = rightFunction(rune(tokens[index][1]))

	tokens[index] = " "
	for i := range number_of_words {
		tokens[index - i - 1] = exec(tokens[index - i - 1])
	}
}
