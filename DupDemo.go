package main

import (
	"bufio"
	"fmt"
	"os"
)

func init() {
	file, err := os.Open("Duptest.in")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	os.Stdin = file
}

func main() {
	cnt := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		cnt[input.Text()]++
	}

	for line, stc := range cnt {
		if stc > 1 {
			fmt.Printf("%d\t%s\n", stc, line)
		}
	}
}
