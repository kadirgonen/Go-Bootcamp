package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	//eşzamnlı yapılan her işlemin bir tanesine goroutine denir. eş zamanlı çaılşmaya concurrency denir
	//paralelism aynı anda yapmak, concurrency tek çekirdekli mantıkda geçerli gibi

	wg.Add(1) //maine senden başka bekiyecek bir go routine daha var

	go printX() //go diyince paralel işleme alıp altdaki diğer go rytinlere de geçebilir.
	//fmt.Println()
	wg.Wait() //x ler daha önce yazılsın istersen
	printY()
	//time.Sleep(time.Second) //performansı düşürür wg kullan wg gelmeden bu olmadan main biter main biter xy basılmaz

}

func printX() {
	for i := 0; i < 1000; i++ {
		fmt.Print("X")
	}
	wg.Done() //wait bitir
}

func printY() {
	for i := 0; i < 20; i++ {
		fmt.Print("Y")
	}
}

// main goroutine
