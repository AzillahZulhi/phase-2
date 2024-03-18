package main

import (
	"fmt"
)

func sendNumbers(evenCh, oddCh chan<- int, errCh chan<- error) {
	defer close(evenCh)
	defer close(oddCh)
	defer close(errCh)

	for i := 1; i <= 25; i++ {
		if i > 20 {
			errCh <- fmt.Errorf("Number %d is greater than 20", i)
			continue
		}

		if i%2 == 0 {
			evenCh <- i
		} else {
			oddCh <- i
		}
	}
}

func printEvenNumbers(evenCh <-chan int, done chan<- bool) {
	for num := range evenCh {
		fmt.Println("Received even number:", num)
	}
	done <- true
}

func printOddNumbers(oddCh <-chan int, done chan<- bool) {
	for num := range oddCh {
		fmt.Println("Received odd number:", num)
	}
	done <- true
}

func main() {
	evenCh := make(chan int)
	oddCh := make(chan int)
	done := make(chan bool)
	errCh := make(chan error)

	go sendNumbers(evenCh, oddCh, errCh)
	go printEvenNumbers(evenCh, done)
	go printOddNumbers(oddCh, done)

	for i := 0; i < 5; i++ {
		select {
		case err := <-errCh:
			fmt.Println("Error:", err)
		case <-done:
		}
	}
}
