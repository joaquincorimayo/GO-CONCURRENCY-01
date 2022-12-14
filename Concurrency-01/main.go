package main

import (
	"fmt"
	"time"
)

func hiNoConcurrency(num int) {
	fmt.Println("Hi, ", num)
	time.Sleep(1000 * time.Millisecond)
}

// PROC
func main() {
	duration := time.Now()

	for i := 0; i < 10; i++ {
		hiNoConcurrency(i)
	}

	fmt.Println("Duration: ", time.Since(duration))
	var closeFinal string
	fmt.Scan(&closeFinal)
}
