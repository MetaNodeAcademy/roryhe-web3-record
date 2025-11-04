package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

/*
题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
考察点 ： go 关键字的使用、协程的并发执行。
*/

func basicUse() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			if i%2 != 0 {
				fmt.Print(i)
			}
		}
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			if i%2 == 0 {
				time.Sleep(1000)
				fmt.Print(i)
			}
		}
	}(&wg)

	wg.Wait()
}

/*
*
题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
考察点 ：协程原理、并发任务调度。
*/
type TaskGoroutine struct {
	id       int
	msg      string
	duration time.Duration
}

func dispatchFunction(wg *sync.WaitGroup, task TaskGoroutine, start time.Time) {
	defer wg.Done()
	task.duration = time.Since(start)
	fmt.Println(task.id, task.msg, task.duration)
}

func main() {
	basicUse()
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		task := TaskGoroutine{
			id:  i,
			msg: strings.Join([]string{"task:", strconv.Itoa(i)}, "-"),
		}
		go dispatchFunction(&wg, task, time.Now())
	}
	wg.Wait()
	fmt.Println(time.Now(), "all task done")
}
