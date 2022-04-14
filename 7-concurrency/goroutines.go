package main

import (
	"fmt"
	"sync"
)

func runPrint(name int)  {
	fmt.Println("Hello from goroutine", name)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(localI int) {
			runPrint(localI)
			wg.Done()
		} (i)
	}
	// the Goroutines can be launched also when the for loop exits
	
	wg.Wait()
}

