package main

import (
	"fmt"
	"errors"
	"os"
)


func divMod(a int, b int) (int, int, error) {
	if b == 0 {
		return 0, 0, errors.New("Division by 0")
	}
	return a / b, a % b, nil
}

func divModME(a int, b int) (int, int, error) {
	if b == 0 {
		return 0, 0, &MyError{A: a, B: b, Message: "Division by 0"}
	}
	return a / b, a % b, nil
}

type MyError struct {
	A int
	B int
	Message string
}

func (me *MyError) Error() string {
	return fmt.Sprintf("value %d and %d produced error %s", me.A, me.B, me.Message)
}

func reallyNil() error {
	var e error
	fmt.Println("e is nil", e == nil)
	return e
}

func notReallyNil() error {
	var me *MyError
	fmt.Println("me is nil:", me == nil)
	return me
}

func main() {
	div, mod, err := divMod(10, 20)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(div, mod)

	div1, mod1, err1 := divModME(10, 20)
	if err1 != nil {
		fmt.Println(err1)
		os.Exit(1)
	}
	fmt.Println(div1, mod1)

	e := reallyNil()
	me := notReallyNil()
	
	fmt.Println("e is nil:", e == nil)
	fmt.Println("me is nil:", me == nil)
	// there is a underlying type associated with that error interface that was returned
}
