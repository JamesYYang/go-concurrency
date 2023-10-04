package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// fan out: 将一个channel的消息发到多个channel
// 也就是一个输入，多个输出

func doTask(name string) chan<- string {
	c := make(chan string)
	go func() {
		for msg := range c {
			fmt.Printf("[%s] %s \n", name, msg)
		}
	}()
	return c
}

func fanOutSimple(out ...chan<- string) chan<- string {
	c := make(chan string)
	go func(ch <-chan string) {
		for msg := range ch {
			for _, co := range out {
				co <- msg
			}
		}
	}(c)
	return c

}

func main() {

	task1 := doTask("do homework")
	task2 := doTask("do housework")

	c := fanOutSimple(task1, task2)

	for i := 0; i < 10; i++ {
		c <- strconv.Itoa(i)
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	}

	fmt.Println("Task done. I'm leaving")
}
