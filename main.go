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

	out_file, out_err := os.OpenFile(os.Args[2], os.O_CREATE | os.O_WRONLY | os.O_TRUNC, 0644)
	if out_err != nil {
		fmt.Println("Error opening/creating the output file")
	}
	defer out_file.Close()

	out_file.WriteString(piscine.Capitalize(string(data)))
}
