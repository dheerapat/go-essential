package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	// ch <- 353

	go func() {
		ch <- 353
	}()

	val := <-ch
	fmt.Printf("got %d\n", val)

	fmt.Println("---")

	const count = 3
	go func() {
		for i := 0; i < count; i++ {
			fmt.Printf("sending %d\n", i)
			ch <- i
			time.Sleep(time.Second)
		}
	}()

	for i := 0; i < count; i++ {
		val := <-ch
		fmt.Printf("receive %d\n", val)
	}

	fmt.Println("---")

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("sending %d\n", i)
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
	}()

	for i := range ch {
		fmt.Printf("receive %d\n", i)
	}

	fmt.Println("---")

	ch2 := make(chan int, 1)
	ch2 <- 19
	val2 := <-ch2
	fmt.Println(val2)
}
