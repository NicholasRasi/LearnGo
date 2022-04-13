package main

import "fmt"

type tester interface {
	test(int) bool
}

func runTests(i int, tests []tester) bool {
	result := true
	for _, test := range tests {
		result = result && test.test(i)
	}
	return result
}

type rangeTest struct {
	min int
	max int
}

func (rt rangeTest) test(i int) bool {
	return rt.min <= i && i <= rt.max
}

type divTest int

func (dt divTest) test(i int) bool {
	return i%int(dt) == 0
}

func doFoo(i interface{}) {
	switch i := i.(type) {
	case int:
		fmt.Println("Double i is", i*2)
	case string:
		fmt.Println("len i is", len(i))
	default:
		fmt.Println("?")
	}
}

type testerFunc func(int) bool

func (tf testerFunc) test(i int) bool {
	return tf(i)
}
// testerFunc implements the tester interface


func main() {
	result := runTests(10, []tester{
		rangeTest{min: 1, max: 10},
		divTest(5),
	})
	fmt.Println(result)

	var i interface{}
	i = "Hello"
	j := i.(string)
	// type assertion
	k, ok := i.(int)
	fmt.Println(j, k, ok)

	// l := i.(int)
	// fmt.Println(m)

	doFoo(10)
	doFoo("Hi")
	doFoo(false)

	result2 := runTests(10, []tester{
		testerFunc(func(i int) bool {
			return i%2 == 0
		}),
		testerFunc(func(i int) bool {
			return i < 20
		}),
	})
	fmt.Println(result2)
}

