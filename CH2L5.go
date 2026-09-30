// Assignment 

// Boots is a bear. (Not a dog, haters).   ---- done

// Run the code as-is. Notice that the simple string "boots" has 5 bytes, and 5 runes (characters).
// Update the name constant to be the bear emoji instead of the word "boots". --- done 


package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const name = "🐻"
	fmt.Printf("constant 'name' byte length: %d\n", len(name))
	fmt.Printf("constant 'name' rune length: %d\n", utf8.RuneCountInString(name))
	fmt.Println("=====================================")
	fmt.Printf("Hi %s, so good to have you back in the arcanum\n", name)
}
