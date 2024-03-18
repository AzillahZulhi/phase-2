package main

import (
	"fmt"
)

func sendNumbers(evenCh, oddCh chan<- int) {
	defer close(evenCh)
	defer close(oddCh)

	for i := 1; i <= 20; i++ {
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

	go sendNumbers(evenCh, oddCh)
	go printEvenNumbers(evenCh, done)
	go printOddNumbers(oddCh, done)

	<-done
	<-done
}
