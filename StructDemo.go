package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {

	fmt.Println(person{"Bob", 20})

	alice := person{name: "Alice", age: 30}

	fmt.Printf("%#v\n", alice)

	fred := person{name: "Fred"}

	fmt.Printf("%+v\n", fred)

	fmt.Println(&person{"Ann", 40})
	fmt.Println(newPerson("Jon"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)

	dog := struct {
		name   string
		isGood bool
	}{
		"Colarie",
		true,
	}
	fmt.Println(dog)
}
