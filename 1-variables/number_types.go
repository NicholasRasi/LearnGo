package main

import "fmt"

func main()  {
	var i int8 = 10
	var j float32 = 1.3
	var l int32 = 20

	// fmt.Println(i + j) does not work
	fmt.Println(float32(i) + j)
	fmt.Println(i + int8(j+1.9))
	
	// fmt.Println(i + l) does not work
	fmt.Println(i + int8(l))
}