# Functions
```go
func main() {
    ...
}
```

- parameters are specified with name and type
- all the parameters are required
- no optional parameters, or named parameters
- no function overloading (same function name but different parameters)
- the return type is required
- function calls are call by value
- functions can not be declared inside another function, but can be declared as anonymous function and assigned to a value
- function can be passed as a parameter
- the value of a local variable can be modified in using an anonymous local function

### Pointers
- &: get the pointer
- *: get the value

Pass parameter as a pointer (reference)
```go
e := 10
setTo100(&e)

func setTo100(a *int) {
	*a = 10
}
```