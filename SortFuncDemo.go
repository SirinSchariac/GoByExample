package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	fruits := []string{"banana", "peach", "kiwi", "dragonfruit"}

	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}
	slices.SortFunc(fruits, lenCmp)
	fmt.Println("Sorted fruits:", fruits)

	type Person struct {
		name string
		age  int
	}

	people := []Person{
		Person{name: "Alice", age: 30},
		Person{name: "Bob", age: 25},
		Person{name: "Charlie", age: 35},
		Person{name: "Dave", age: 20},
	}
	slices.SortFunc(people,
		func(a, b Person) int {
			return cmp.Compare(a.age, b.age)
		})
	fmt.Println("Sorted people by age:", people)
}
