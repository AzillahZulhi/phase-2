package main

import (
	"fmt"
	"sync"
)

func producer(ch chan<- int, wg *sync.WaitGroup) {
	defer close(ch)
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	wg.Done()
}

func consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range ch {
		fmt.Println(num)
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int, 5) // Modified Channel

	wg.Add(2)
	go producer(ch, &wg)
	go consumer(ch, &wg)

	wg.Wait()
}

// Saluran Tanpa Buffer:
// - Saat menggunakan saluran tanpa buffer, pengiriman data akan berhenti sampai ada penerima yang siap untuk membacanya.
// - Proses pengiriman dan penerimaan data berlangsung secara sinkron.

// Saluran dengan Buffer:
// - Saluran dengan buffer memungkinkan untuk mengirim data ke dalam buffer tanpa menunggu penerima membaca data tersebut.
// - Jika buffer belum penuh, pengirim dapat mengirim data tanpa harus menunggu.
// - Jika buffer sudah penuh, pengiriman data akan terhenti sementara sampai ada ruang kosong di dalam buffer.
// - Proses pengiriman dan penerimaan data berlangsung secara asinkron.
