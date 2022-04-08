package main

import "fmt"

func main() {
	var i int = 10
	fmt.Println(i)

	l := 0
	fmt.Println(l)

	var j int
	fmt.Println(j)

	j = 10
	fmt.Println(j)

	var s string
	s = "Hello, world!"
	fmt.Println(s)

	s = `Hello,
	world!`
	fmt.Println(s)

	fmt.Println(string(s[0]))
	fmt.Println(string(s[:5]))

	r := '♥'
	fmt.Println(string(r))

	var vals [3]int
	vals[0] = 2
	fmt.Println(vals)
}
