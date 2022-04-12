package main

import "fmt"

func main()  {
	m := make(map[string]int)
	m["hi"] = 1

	fmt.Println(m["a"])

	if v, ok := m["hi"]; ok {
		fmt.Println("hi key in m:", v)
	}

	m2 := map[string]int {
		"a": 1,
		"b": 10,
		"c": 100,
	}

	for k, v := range m2 {
		fmt.Println(k, v)
	}

	delete(m2, "b")

	fmt.Println(m2)
}