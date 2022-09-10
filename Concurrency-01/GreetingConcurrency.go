package main

import (
	"fmt"
	"time"
)

func hiConcurrency(num int) {
	fmt.Println("Hi, ", num)
	time.Sleep(1000 * time.Millisecond)
}

func main() {
	duration := time.Now()

	// Using go routines
	for i := 0; i < 10; i++ {
		go hiConcurrency(i) // creates 10 processes [i]
	}

	fmt.Println("Duration: ", time.Since(duration))
	var closeFinal string
	fmt.Scan(&closeFinal)
}
