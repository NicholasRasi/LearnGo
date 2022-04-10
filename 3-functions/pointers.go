package main

import "fmt"

func main() {
	a := 10
	b := &a
	c := a

	fmt.Println(a, b, *b, c)

	a = 20
	fmt.Println(a, b, *b, c)

	*b = 30
	fmt.Println(a, b, *b, c)

	c = 40
	fmt.Println(a, b, *b, c)

	d := new(int)
	fmt.Println(d)
	fmt.Println(*d)

	e := 20
	fmt.Println(a)
	setTo100(&e)
	fmt.Println(e)
}

func setTo100(a *int) {
	*a = 10
}
