package main

import (
	"fmt"
	"time"
)

/*
实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
考察点 ：通道的缓冲机制。
*/

func producer(ch chan int) {
	for i := 0; i < 100; i++ {
		ch <- i
		fmt.Println("producer ", i)
		time.Sleep(20 * time.Millisecond)
	}
	close(ch)
}

func consumer(ch chan int) {
	for c := range ch {
		fmt.Println("consumer ", c)
		time.Sleep(50 * time.Millisecond)
	}
}

func main() {
	ch := make(chan int, 3)
	go producer(ch)
	go consumer(ch)
	time.Sleep(5 * time.Second)
	fmt.Println("main end")
}
