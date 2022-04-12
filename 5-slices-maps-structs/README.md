# Slices, Maps and Structs
### Slices
> A slice is a growable sequence of values of a single specified type

```go
s := make([]string, 0)
s = append(s, "hello")

s1 := []string{"a", "b", "c"}
```

- slices are reference types
```go
s1 := []string{"a", "b", "c"}

s2 := s1[0:2]
s1[0] = "d"
fmt.Println(s2) // d, b
```

### Maps
> A map is a data type that associates a value of one data type to a value of
another data type

- (Java hashmap, Python dict, Javascript Object, Ruby hash)
- key cannot be: slice, map, function, type that contains slice, map or function
```go
m := make(map[string]int)
m["hi"] = 1
```
- to check if a value is in the map:
```go
if v, ok := m["hi"]; ok {
    fmt.Println("hi key in m:", v)
}
```
- iterate a map:
```go
for k, v := range m2 {
    fmt.Println(k, v)
}
```
- delete a value with ``delete(m, "hello")```
- maps are reference types

### Structs
- structs are not objects
- no inheritance for structs

```go
type Foo struct {
	A int
	B string
}
type Baz struct {
	D string
	Foo // embedded struct
}

f := Foo{}

f1 := Foo{1, "Hello"} // not common

f2 := Foo{
    b: "Hi",
}

f2.A = 100

b := Baz{D: "Hi", Foo: f1}
```

- struct fields whose names start with capital letters are visible outside of their package
- struct fields whose names start with capital letters are also visible outside of their packages
- structs are value type, so they are copied
