package main

import (
	"errors"
	"fmt"
)

type argError struct {
	arg int
	msg string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.msg)
}

func f(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg: arg, msg: "Can't handle it"}
	}
	return arg + 3, nil
}

func main() {
	_, err := f(42)
	var ae *argError
	if errors.As(err, &ae) {
		fmt.Println("Error:", ae.arg, " ", ae.msg)
	} else {
		fmt.Println("err does not match argError")
	}
}
