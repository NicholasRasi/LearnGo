package main

import "fmt"

func runPrint(name int)  {
	fmt.Println("Hello from goroutine", name)
}

func main() {
	in := make(chan string)
	out := make(chan string)
	go func() {
		// read from in
		name := <-in
		// write to out
		out <- fmt.Sprintf("Hello, " + name)
	}()
	in <- "Jane"
	// close the channel
	close(in)
	message := <-out
	fmt.Println(message)


	outbuf := make(chan int, 10)
	for i := 0; i < 10; i++ {
		go func (localI int) {
			outbuf <- localI * 2
		}(i)
	}
	var result []int
	for i:= 0; i < 10; i++ {
		val := <-outbuf
		result = append(result, val)
	}
	fmt.Println(result)
}

