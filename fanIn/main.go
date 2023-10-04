package main

import (
	"fmt"
	"math/rand"
	"time"
)

// fan in: 将多个channel的消息合并到一个channel
// 也就是多个输入，一个输出

func doTask(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		}
	}()
	return c // return a channel to caller.
}

func fanInSimple(cs ...<-chan string) <-chan string {
	c := make(chan string)
	for _, ci := range cs { // spawn channel based on the number of input channel
		go func(cv <-chan string) { // cv is a channel value
			for {
				c <- <-cv
			}
		}(ci) // send each channel to
	}
	return c
}

func main() {

	task1 := doTask("do homework")
	task2 := doTask("do housework")

	c := fanInSimple(task1, task2)

	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}

	fmt.Println("Task done. I'm leaving")
}
