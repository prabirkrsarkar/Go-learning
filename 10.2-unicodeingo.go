package main

import (
	"fmt"
	"unicode"
)

func main() {

	// Indexing a Go string returns bytes, not Unicode characters (runes).
	// UTF-8 characters may occupy multiple bytes, so iterating byte-by-byte
	// can produce incorrect character output for non-ASCII text.

	word := "バーラト"
	fmt.Println("-------Using a for loop and string indexing:--------")
	for i := 0; i < len(word); i++ {
		fmt.Printf("Character : %c, Type : %T\n", word[i], word[i])
	}

	// To correctly iterate over Unicode characters, use a range loop
	// Each iteration gives you the index and the rune (Unicode code point)
	// unicode.IsLetter() expects a rune as an argument
	fmt.Println("-------Using a range loop to get Unicode characters:--------")
	for _, r := range "バーラト" {
		fmt.Printf("Character : %c, Type : %T, Is Letter: %t, Is Graphic: %t\n", r, r, unicode.IsLetter(r), unicode.IsGraphic(r))
	}
}
