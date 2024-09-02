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

// modified input text hexadecimal number format 1a -> 1A
func TestTwo(t *testing.T) {
	inFile, inErr := os.OpenFile("sample.txt", os.O_CREATE | os.O_WRONLY | os.O_TRUNC, 0644)
	if inErr != nil {
		t.Fatal("Error creating sample.txt file")
	}
	inFile.WriteString("I have to pack 101 (bin) outfits. Packed 1A (hex) just to be sure")
	goReloaded("sample.txt", "result.txt")
	outFile, outErr := os.ReadFile("result.txt")
	if outErr != nil {
		t.Fatal("Error reading from result.txt file")
	}
	if string(outFile) != "I have to pack 5 outfits. Packed 26 just to be sure" {
		t.Fatal("Failed. result.txt should be instead:\nI have to pack 5 outfits. Packed 26 just to be sure")
	}
}

func TestThree(t *testing.T) {
	inFile, inErr := os.OpenFile("sample.txt", os.O_CREATE | os.O_WRONLY | os.O_TRUNC, 0644)
	if inErr != nil {
		t.Fatal("Error creating sample.txt file")
	}
	inFile.WriteString("Don not be sad ,because sad backwards is das . And das not good")
	goReloaded("sample.txt", "result.txt")
	outFile, outErr := os.ReadFile("result.txt")
	if outErr != nil {
		t.Fatal("Error reading from result.txt file")
	}
	if string(outFile) != "Don not be sad, because sad backwards is das. And das not good" {
		t.Fatal("Failed. result.txt should be instead:\nDon not be sad, because sad backwards is das. And das not good")
	}
}

func TestFour(t *testing.T) {
	inFile, inErr := os.OpenFile("sample.txt", os.O_CREATE | os.O_WRONLY | os.O_TRUNC, 0644)
	if inErr != nil {
		t.Fatal("Error creating sample.txt file")
	}
	inFile.WriteString("harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '")
	goReloaded("sample.txt", "result.txt")
	outFile, outErr := os.ReadFile("result.txt")
	if outErr != nil {
		t.Fatal("Error reading from result.txt file")
	}
	if string(outFile) != "Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'" {
		t.Fatal("Failed. result.txt should be instead:\nHarold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'")
	}
}