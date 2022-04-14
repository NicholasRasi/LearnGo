# Concurrency
> Concurrency in Go is based on CSP and built into the language

### Goroutines
- CSP stands for communicating sequential processes
- goroutines are lightweight processes managed by the Go runtime, faster to create, efficient with memory and faster to switch
- in a single process tens of thousands goroutines can coexist
- goroutines are managed by the Go runtime scheduler that creates a number of threads and then schedules goroutines across threads automatically
- a Goroutine is launched with ``go``
```go
go runPrint()
```
- it is possible to wait until a Goroutine finishes its work with a wait group

### Channels
- a type in Go for transferring data between goroutines
- one or more goroutines write to a channel and one or more goroutines read back from the same channel
- data on channels is type and by default channel reads ad writes are synchronous
- if the passed value is a value type (int, string) a copy of the data is made, otherwise it the value is a reference type (pointer, map, slice) a pointer to the data is put on the channel
- create a channel and read/write from/to:
```go
in := make(chan string)
name := <-in
out <- fmt.Sprintf("Hello, " + name)
```
- a buffered channel allows to write a specified number of times before there's a read on the channel, it is not infinite
```go
outbuf := make(chan int, 10)
```
- it is possible to check if a channel is opened with this idiom
```go
i, ok := <-in
if !ok {
    close(out)
    break
}
```
- channels can be iterated:
```go
for v := range out {
    // use v
}
```
- channels are reference types, so the zero value for a channel is nil

|  | unbuffered | buffered | nil | closed |
|---|---|---|---|---|
| read | Pause until something is written | Pause if buffer is empty | Hang forever | Return immediately with zero value (use comma-ok to see if it is closed) |
| write | Pause until something is read | Pause only if buffer full | Hang forever | Panic |
| close | Ok | Ok | Panic | Panic |

### Select
> Select statements make concurrency in Go interesting

```go
select {
case in <- 2:
	fmt.Println("wrote 2 to in")
case v := <-out:
	fmt.Println("read", v, "from out")
default:
}
```
- if more than one case in the select statement can be run, because multiple channels are ready to be read or written, the select statement picks one case at random to run
- the default in select will never run if any case can be run instead
- if you ever want to disable a case in the select, set the channel variable to nil and the case will never succeed
- Go automatically manages memory with garbage collection but it does not automatically manage goroutines. Anytime you launch a goroutine, you must provide a way for it to exit otherwise resources are wasted. The done channel pattern is used to prevent resource wasting
- backpressure is used to ignore requests that can not be handled due to capacity