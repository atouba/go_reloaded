package main

import (
	"fmt"
	"os"
	"github.com/atouba/piscine"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run <arg 1> <arg 2>")
		return
	}

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error reading input file")
	}

	out_file, out_err := os.OpenFile(os.Args[2], os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0644)
	if out_err != nil {
		fmt.Println("Error opening/creating the output file")
	}
	defer out_file.Close()

	tokens := tokenize(string(data))
	applyChanges(tokens)
	writeOutput(tokens, out_file)
}

func writeOutput(tokens []string, out_file *os.File) {
	//  a word is considered an actual word preceded by a space
	//  and appended by a punctituation mark if there's any
	var word string
	for i := 0; i < len(tokens); i++ {
		if tokens[i][0] == ' ' {
			continue
		}
		word = tokens[i]
		if !isPunctuation(rune(tokens[i][0])) && i != 0 {
			word = piscine.Concat(" ", word)
		}
		if i < len(tokens) - 1 && isPunctuation(rune(tokens[i + 1][0])) {  // how come tokens[][0] is not rune
			word = piscine.Concat(word, tokens[i + 1])
			tokens[i + 1] = " "
		}
		out_file.WriteString(word)
		// out_file.WriteString(" ")
	}
}
