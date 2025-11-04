package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

/**
题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
考察点 ：通道的基本使用、协程间通信。
*/

type Task struct {
	id       int
	msg      string
	duration time.Duration
}

func send(channel chan<- Task, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(channel)
	for i := 1; i <= 10; i++ {
		task := Task{
			id:  i,
			msg: strings.Join([]string{"hello", strconv.Itoa(i)}, "-"),
		}
		channel <- task
		fmt.Println("send ", task.id, task.msg)
		time.Sleep(time.Second)
	}
}

func receive(channel <-chan Task, wg *sync.WaitGroup) {
	defer wg.Done()
	for ch := range channel {
		now := time.Now()
		time.Sleep(time.Duration(200+ch.id*30) * time.Millisecond)
		ch.duration = time.Since(now)
		fmt.Println("receive:", ch.id, ch.msg, ch.duration)
	}
}

func main() {
	var wg sync.WaitGroup
	var channel chan Task = make(chan Task, 3)
	wg.Add(2)
	go send(channel, &wg)
	go receive(channel, &wg)
	wg.Wait()
	fmt.Println("all done")
}
