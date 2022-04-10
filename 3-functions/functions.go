package main

import "fmt"

func main() {
	fmt.Println(addNumbers(1, 2))

	_, remainder := divAndRemainder(100, 2)
	fmt.Println(remainder)

	myAddNumbers := addNumbers
	fmt.Println(myAddNumbers(10, 20))

	b := 10
	// closure
	myFunc := func(a int) int {
		b = 20
		return a + b
	}
	fmt.Println(myFunc(10))

	printOperation(10, addOne)
}

func addNumbers(num1 int, num2 int) int {
	return num1 + num2
}

func divAndRemainder(a int, b int) (int, int) {
	return a / b, a % b
}

func addOne(a int) int {
	return a + 1
}

func addTwo(a int) int {
	return a + 2
}

func printOperation(a int, f func(int)int) {
	fmt.Println(f(a))
}

