package main

import (
	"bufio"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func getRequestNumber(numberRequest int) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/" + strconv.Itoa(numberRequest))

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	fmt.Println("Connection closed | Status: ", resp.Status)
	// Show me request.
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan(); i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

}

func main() {
	duration := time.Now()

	for i := 0; i < 100; i++ {
		go getRequestNumber(i)
	}

	fmt.Println("Duration: ", time.Since(duration))
	var closeFinal string
	fmt.Scan(&closeFinal)
}
