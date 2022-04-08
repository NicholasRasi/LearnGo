# Variables

## Declaration
- ``var <name> <type> = <init_value>``
- Go allows to take advantage to type inference: ``var i = 10``
- Short syntax: ``i := 10``
    - it can not be used outside of a function
    - it does not allow to declare the variable type
- all declared variables have a value
- declared variables must be used (read), otherwise a compilation error is triggered

## Built-in Types

### Numeric Types
- signed integer: int8 ... int64
- unsigned integer: uint8 ... uint64
- floating point: float32, float64
- boolean

> Go has all the usual math operators

#### Special Types
- byte is same as uint8
- int is same as int32 or int64 (depending on the size of a native int on your computer)
- uint is same as uint32 or uint64 (depending on the size of a native int on your computer)
- rune is same as int32

> Go won't automatically convert types

#### Type Conversion
- int8()
- 0 does not convert to boolean False

### String Type
- interpreted literal
    - double quotes to mark the beginning and end of the string
    - ``\`` to escape ``"``
    - ``\n\t`` newline, tab
- raw string literal
    - backtick
    - can go multiple lines

#### String Operations
- concat: plus operator ``s1 + s2``
- substring: ``s[<start>:<end>]``
- string len: ``len(s)``


### Array
- ``var <name> [<size>]<type>``
- the length of the array is considered part of the type
- there is no type conversion, e.g. is not possible to do
```
var vals [3]int
var vals2 [2]int = vals
```
- slice is used and more versatile