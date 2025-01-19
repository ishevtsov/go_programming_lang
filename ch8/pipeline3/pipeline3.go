package main

import (
	"fmt"
	"time"
)

// Counter
func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

// Squarer
func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

// Printer
func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
		time.Sleep(25 * time.Millisecond)
	}
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}
