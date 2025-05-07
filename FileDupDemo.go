package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	cnt := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "FileDupDemo: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			if line != "" {
				cnt[line]++
			}
		}
	}
	for line, n := range cnt {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
