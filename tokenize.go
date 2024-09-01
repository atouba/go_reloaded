package main

import (
	"github.com/atouba/piscine"
)

func isPunctuation(c rune) bool {
	if c == '.' || c == ',' || c == '!' ||
	c == '?' || c == ':' || c == ';' ||
	c == '\'' || c == '(' {
		return true
	}
	return false
}

func token(data string) string {
	i := 0
	for ; i < piscine.StrLen(data);  {
		if data[0] == '(' && data[i] == ')' {
			i++
			break
		} else if (data[i] == ' ' || isPunctuation(rune(data[i]))) && data[0] != '(' {
			break
		}
		i++
	}
	if i == 0 {
		i++
	}
	return data[:i]
}

func tokenize(data string) (tokens []string) {
	for i := 0; i < piscine.StrLen(data); {
		for data[i] == ' ' {
			i++
		}
		tokens = append(tokens, token(data[i:]))
		i += piscine.StrLen(tokens[len(tokens) - 1])
	}
	return tokens
}
