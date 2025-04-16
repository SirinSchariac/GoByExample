package main

import (
	"fmt"
	"unicode/utf8"
)

// main demonstrates the handling of strings and runes in Go. It includes:
// - Calculating the length of a string in bytes using len().
// - Iterating over a string byte by byte and printing each byte in hexadecimal format.
// - Counting the number of runes in a string using utf8.RuneCountInString().
// - Iterating over a string using a range loop to extract runes and their starting positions.
// - Using utf8.DecodeRuneInString() to decode runes manually and determine their widths.
//
// The difference between %U and %+U in fmt.Printf formatting is:
// - %U: Prints the Unicode code point of the rune in the format "U+XXXX".
// - %+U: Prints the Unicode code point along with its Go syntax representation, e.g., 'A' for U+0041.
func main() {
	//const s = "こんにちは, 世界"
	const s = "Hello world"
	fmt.Println("Len of s:", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("% x", s[i])
	}
	fmt.Println()

	fmt.Println("Rune count: ", utf8.RuneCountInString(s))

	for idx, rune_val := range s {
		fmt.Printf("%#U starts at %d\n", rune_val, idx)
	}
	fmt.Println("\nUsing DecodeRuneInString:")
	for i, w := 0, 0; i < len(s); i += w {
		rune_val, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%+U starts at %d\n", rune_val, i)
		w = width
	}
}
