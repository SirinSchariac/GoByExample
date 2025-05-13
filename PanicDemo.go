package main

import "os"

func main() {
	panic("A problem has occurred")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
