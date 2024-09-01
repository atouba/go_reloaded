package main

import (
	"fmt"
	"os"
)

func goReloaded(in_file_name, out_file_name string) {
	data, err := os.ReadFile(in_file_name)
	if err != nil {
		fmt.Println("Error reading input file")
		os.Exit(1)
	}

	out_file, out_err := os.OpenFile(out_file_name, os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0644)
	if out_err != nil {
		fmt.Println("Error opening/creating the output file")
		// is it okay to exit without closing the file?
		// also for the output is the "exit status " output okay?
		os.Exit(1)
	}
	defer out_file.Close()

	tokens := tokenize(string(data))
	applyChanges(tokens)
	writeOutput(tokens, out_file)
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run <arg 1> <arg 2>")
		return
	}

	goReloaded(os.Args[1], os.Args[2])
}
