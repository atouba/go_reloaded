package main

import (
	"github.com/atouba/piscine"
)

func token(data string) string {
	i := 0
	for ; i < piscine.StrLen(data) && data[i] != ' '; {
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