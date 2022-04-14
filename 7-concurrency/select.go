package main

import "fmt"

func multiples(i int) (chan int, chan struct{}) {
	out := make(chan int)
	done := make(chan struct{})
	curVal := 0
	go func() {
		for {
			select {
			case out <- curVal * i:
				curVal++
			case <-done:
				fmt.Println("goroutine shutting down")
				return
			}
		}
	}()
	return out, done
}

func main() {
	in := make(chan int, 1)
	out := make(chan int, 1)

	out <- 1

	select {
	case in <- 2:
		fmt.Println("wrote 2 to in")
	case v := <-out:
		fmt.Println("read", v, "from out")
	default:
		fmt.Println("?")
	}


	twosCh, done := multiples(2)
	for v := range twosCh {
		if v > 20 {
			break
		}
		fmt.Println(v)
	}
	close(done)
}

