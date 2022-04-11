package main

import (
	"fmt"
	"strings"
	"math"
	"math/rand"
	"time"
	add "4-packages/adder"
)

func add1(in rune) rune {
	return (((in - 'a') + 1) % 26) + 'a'
}

func main()  {
	s := "test"
	s2 := strings.Map(add1, s)
	fmt.Println(s2)

	rand.Seed(time.Now().UnixNano())
	a := rand.Intn(10)
	b := rand.Intn(10)
	c := int(math.Max(float64(a), float64(b)))
	fmt.Println(a, b, c)

	fmt.Println(add.Addnum(10))
}