package main

import (
	"errors"
	"fmt"
)

func f(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("Can't handle 42")
	}

	return arg + 3, nil
}

var ErrOutOfTea = fmt.Errorf("out of tea")
var ErrPower = fmt.Errorf("Can't boil water")

func makeTea(arg int) error {
	if arg == 2 {
		return ErrOutOfTea
	} else if arg == 4 {
		return fmt.Errorf("Making tea: %w", ErrPower)
	}
	return nil
}

func main() {
	for _, i := range []int{7, 42} {

		if r, e := f(i); e != nil {
			fmt.Println("f failed: ", e)
		} else {
			fmt.Println("f worked: ", r)
		}
	}

	for i := range 5 {
		if err := makeTea(i); err != nil {
			if errors.Is(err, ErrOutOfTea) {
				fmt.Println("We need to buy new tea")
			} else if errors.Is(err, ErrPower) {
				fmt.Println("We need to pay the power bill")
			} else {
				fmt.Printf("Unknown error: %v\n", err)
			}
			continue
		}
		fmt.Println("Tea is ready")
	}
}
