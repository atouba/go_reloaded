package main

import "github.com/atouba/piscine"

func isVowelStart(word string) bool {
	vowels := []rune{'a', 'e', 'i', 'o', 'h'}
	for _, v := range vowels {
		if rune(word[0]) == rune(v) || rune(word[0]) == rune(v - 32) {     // !?
			return true
		}
	}
	return false
}

// a -> an
// (hex) (bin)   (up,) (low,) (cap)
// punctiuation . , ...    ' last step when concatinating
func applyChanges(tokens []string) {
	for i, token := range tokens {
		if token[0] == '(' {
			applyCaseNumericConversion(tokens, i)
		} else if i + 1 <= len(tokens) &&
		(token == "a" || token == "A") && isVowelStart(tokens[i + 1]) {
			token = token + "n"
		}
	}
	// return tokens
}

func applyCaseNumericConversion(tokens []string, index int) {
	// convertCasefuncs := []func(string) (string) {piscine.ToUpper, piscine.ToLower, piscine.Capitalize}
	// var exec func (string) (string)
	// switch tokens[index][1] {
	// case 'l':
	// 	exec = convertCasefuncs[1]
	// case 'u' :
	// 	exec = convertCasefuncs[0]
	// case 'c' :
	// 	exec = convertCasefuncs[2]
	// }
	// if (index == 0 || index == len(tokens)) {
	// 	return
	// }
	tokens[index] = " "
	// tokens[index - 1] = exec(tokens[index - 1])
	tokens[index - 1] = piscine.ToUpper(tokens[index - 1])
}