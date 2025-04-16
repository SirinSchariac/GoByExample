package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "こんにちは, 世界"
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
		fmt.Printf("%#U starts at %d\n", rune_val, i)
		w = width
	}
}
