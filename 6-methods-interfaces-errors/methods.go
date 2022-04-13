package main

import "fmt"

type Foo struct {
	A int
	B string
}

type Bar struct {
	C bool
	Foo
}

func (f Foo) String() string {
	return fmt.Sprintf("A: %d, B: %s, count: %d", f.A, f.B, f.fieldCount())
}

func (f *Foo) Double() {
	f.A = f.A * 2
}

func (f Foo) fieldCount() int {
	return 2
}

func (b Bar) String() string {
	return fmt.Sprintf("C: %v, f: %s", b.C, b.Foo.String())
}

func (b Bar) fieldCount() int {
	return 3
}

type myInt int

func (mi myInt) isEven() bool {
	return mi%2 == 0
}

func (mi *myInt) Double() {
	*mi = *mi * 2
}

func main() {
	f := Foo {
		A: 1,
		B: "Hello",
	}

	fmt.Println(f.String())
	f.Double()
	fmt.Println(f.String())

	b := Bar {
		C: true,
		Foo: f,
	}
	fmt.Println(b.String())
	b.Double()
	fmt.Println(b.String())

	m := myInt(10)
	fmt.Println(m)
	fmt.Println(m.isEven())
	m.Double()
	fmt.Println(m)
}