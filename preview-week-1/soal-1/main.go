package main

import (
	"fmt"
	"sync"
)

func printNumbers(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}

func printLetters(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 'a'; i <= 'j'; i++ {
		fmt.Printf("%c ", i)
	}
	fmt.Println()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go printNumbers(&wg)
	go printLetters(&wg)

	wg.Wait()
}
