package main

import (
	"os"
	"github.com/atouba/piscine"
)

// check if a token that start with a ' is an actual word (i.e len > 0)
// and that it's not at the start of the line
// so as to determine whether to add a space before the token (word) or not
func isNonStarterWordWithAQuote(tokens []string, word_index int) bool {
	// length of the word preceding the word in question
	// needed to see if its last character is a newline
	l_pre_word := len(tokens[word_index - 1])
	if tokens[word_index][0] == '\'' && len(tokens[word_index]) > 1 &&
	word_index > 0 && tokens[word_index - 1][l_pre_word - 1] != '\n' {
		return true
	}
	return false
}

func writeOutput(tokens []string, out_file *os.File) {
	//  a word is considered an actual word preceded by a space  and appended by a punctituation mark if there's any except in the case of ' where there is 2 cases
	var word string
	for i := 0; i < len(tokens); i++ {
		if tokens[i][0] == ' ' { continue }
		word = tokens[i]
		if i != 0 && (!isPunctuation(rune(tokens[i][0])) ||
		isNonStarterWordWithAQuote(tokens, i)) {
			word = " " + word
		}
		if i + 1 < len(tokens) && isPunctuation(rune(tokens[i + 1][0])) {
			if tokens[i + 1][0] == '\'' && quoteCounting(tokens, i + 1) % 2 == 0 {
				tokens[i + 2] = tokens[i + 1] + tokens[i + 2]
			} else {
				word = piscine.Concat(word, tokens[i + 1])
			}
			tokens[i + 1] = " "
		} else if i + 1 < len(tokens) && i == 0 && rune(tokens[i][0]) == '\'' {
			// for cases like ' a
			word += tokens[i + 1]
			tokens[i + 1] = " "
		}
		out_file.WriteString(word)
	}
}

func quoteCounting(tokens []string, index_word int) int {
	count := 0

	for i := index_word; i < len(tokens); i++ {
		if tokens[i] == "'" {
			count++
		}
	}
	return count
}
