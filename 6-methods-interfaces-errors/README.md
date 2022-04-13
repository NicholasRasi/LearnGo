# Methods, Interfaces, Errors

### Methods
> Go does not have inheritance, but it does have methods

- method with value receiver (it does not modify the value)
```go
func (f Foo) String() string {
	return fmt.Sprintf("A: %d, B: %s", f.A, f.B)
}
```

- method with reference receiver (it modifies the value)
```go
func (f *Foo) Double() {
	f.A = f.A * 2
}
```

- (f Foo) it like the ``self`` first parameter in a Python method declaration or ``this`` in Java
- the compiler won't allow you to call a reference receiver if the struct hasn't been assigned to a variable
- if you call a value receiver on a pointer that's nil, your program will compile, but it will panic at runtime
- no methods overwriting and no virtual method dispatch

### Interfaces
> A type that lists a set of methods but provides no implementation

```go
type tester interface {
	test(int) bool
}
```
- there is no explicit declaration of interface implementation. As long as the type implements the right method (name and parameters), it automatically meets the interface
- empty interface: an interface that has no method at all. It is like the void pointer in C or the object type in Java. It means that it could be anything
```go
var i interface{}
i = "Hi"
// type assertion
k, ok := i.(int)
```
- ``ok`` will be assigned the value of true if the assertion worked and false if it didn't. If the assertion worked, the first variable specified will get the value from the interface. If the assertion fails, it will be assigned a 0 value of that type
- it is possible to have a function that implement an interface by defining a function type and a method on the function type


### Errors
> Go does not have exceptions

-  Go uses a return value to indicate whether or nota function or method completed correctly

```go
type error interface {
	Error() string
}
```

```go
func divMod(a int, b int) (int, int, error) {
	if b == 0 {
		return 0, 0, errors.New("Division by 0")
	}
	return a / b, a % b, nil
}
```