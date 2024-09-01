package main

import (
	"testing"
	"os"
)

func TestOne(t *testing.T) {
	inFile, inErr := os.OpenFile("sample.txt", os.O_CREATE | os.O_WRONLY | os.O_TRUNC, 0644)
	if inErr != nil {
		t.Fatal("Error creating sample.txt file")
	}
	inFile.WriteString("If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?")
	goReloaded("sample.txt", "result.txt")
	outFile, outErr := os.ReadFile("result.txt")
	if outErr != nil {
		t.Fatal("Error reading from result.txt file")
	}
	if string(outFile) != "If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?" {
		t.Fatal("Failed. result.txt should be instead:\nIf I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?")
	}
}
