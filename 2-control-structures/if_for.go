package main

import "fmt"

func main() {
	a := 10
	if a > 5 {
		b := a / 2
		fmt.Println("a is bigger than 5: ", a, b)
	} else {
		// fmt.Println("a is lower or equal than 5: ", a, b) but b is undefined here
	}

	if b := a / 2; a > 5 {
		fmt.Println("a is bigger than 5: ", a, b)
	} else {
		fmt.Println("a is lower or equal than 5: ", a, b)
	}
	// fmt.Println(b) but b is undefined here


	for i := 0; i < 10; i++ {
		if i > 7 {
			break
		}
		fmt.Println(i)
	}

	i := 0
	for i < 10 {
		fmt.Println(i)
		i = i + 1
	}

	j := 0
	for {
		fmt.Println(j)
		j = j + 1
		if j > 10 {
			break
		}
	}

	s := "Hello, world!"
	for k, v := range s {
		fmt.Println(k, v, string(v))
	}
}
