package main

import (
	"encoding/json"
	"fmt"
)

type Foo struct {
	A int
	b string
}

type Bar struct {
	C string
	F Foo
}

// Baz is not a subtype of Foo, it contains a Foo as a field (delegation)
type Baz struct {
	D string
	Foo // embedded struct
}

type Person struct {
	Name string `json:"name"`
	Age	 int	`json:"age"`
}

func main()  {
	f := Foo{}
	fmt.Println(f)

	f1 := Foo{1, "Hello"}
	fmt.Println(f1)

	f2 := Foo{
		b: "Hi",
	}
	fmt.Println(f2)

	f2.A = 100
	fmt.Println(f2)

	b1 := Bar{C: "Goodbye", F: f1}
	fmt.Println(b1)

	bz := Baz{D: "Bye", Foo: f1}
	fmt.Println(bz)
	fmt.Println(bz.Foo)

	jason := `{"name": "Jason", "age": 99}`
	var j Person
	json.Unmarshal([]byte(jason), &j)
	fmt.Println(j)

	j2, _ := json.Marshal(j)
	fmt.Println(string(j2))
}