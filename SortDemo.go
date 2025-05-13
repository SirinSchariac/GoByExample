package main

import (
	"fmt"
	"slices"
)

func main() {
	strs := []string{"c", "b", "a"}
	slices.Sort(strs)
	fmt.Println("Strings:", strs)

	nums := []int{7, 2, 4, 10, 3}
	slices.Sort(nums)
	fmt.Println("Numbers:", nums)

	flag := slices.IsSorted(strs)
	fmt.Println("Is sorted:", flag)

}
