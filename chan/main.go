package main

import (
	"fmt"
	"math/rand"
	"time"
)

// doTask 内部启动任务，返回chan，这样外层可以接受任务返回的消息
func doTask(msg string) <-chan string {
	c := make(chan string)
	go func() {
		// The for loop simulate the infinite sender.
		for i := 0; i < 10; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
		close(c)
	}()
	return c // return a channel to caller.
}

func main() {

	task1 := doTask("do homework")
	task2 := doTask("do housework")

	for i := 0; i < 10; i++ {
		fmt.Println(<-task1)
		fmt.Println(<-task2)
	}

	fmt.Println("Task done. I'm leaving")
}
