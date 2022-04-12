package main

import "fmt"

func main()  {
	s := make([]string, 0)
	fmt.Println(len(s))

	s = append(s, "hello")
	fmt.Println(s)

	s2 := make([]string, 2)
	fmt.Println(s2)
	s2 = append(s2, "hi")
	fmt.Println(s2)

	s3 := []string{"a", "b", "c"}
	for k, v := range s3 {
		fmt.Println(k, v)
	}
	s3 = append(s3, "d")
	fmt.Println(s3)


	s4 := []string{"a", "b", "c"}

	s5 := s4[0:2]
	s4[0] = "d"
	fmt.Println(s5)
}