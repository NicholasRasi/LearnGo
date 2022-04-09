# Control Structures, Functions and Pointers
### If/else Statement
```go
if a > 5 {
    fmt.Println("a is bigger than 5: ", a)
} else {
	fmt.Println("a is lower or equal than 5: ", a)
}
```
- 0 != false

### For Loop
```go
for i := 0; i < 10; i++ {
	fmt.Println(i)
}
```

- no while statement but:

```go
i := 0
for i < 10 {
    fmt.Println(i)
	i = i + 1
}
```

or

```go
j := 0
for {
    fmt.Println(j)
	j = j + 1
    if j > 10 {
        break
    }
}
```

or

```go
s := "Hello, world!"
for k, v := range s {
    fmt.Println(k, v, string(v))
    // k is the offset and v the character (rune)
}
```

### Switch Statement
```go
switch word {
    case "hi":
	    fmt.Println("Very informal!")
		fallthrough
	case "hello":
		fmt.Println("Hi!")
	case "ok":
	case "goodbye", "bye":
		fmt.Println("See you soon!")
	default:
		fmt.Println("What?")
}
```