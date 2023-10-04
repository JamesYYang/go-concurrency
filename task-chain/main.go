package main

import (
	"fmt"
	"math/rand"
	"time"
)

func pass(left, right chan int) {
	v := <-right
	next := v + 1
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	fmt.Printf("pass value %d to next\n", next)
	left <- next
}

func main() {
	const n = 30
	leftmost := make(chan int)
	left := leftmost
	right := leftmost

	for i := 0; i < n; i++ {
		right = make(chan int)
		go pass(left, right)
		left = right
	}

	go func(c chan int) { c <- 1 }(right)
	fmt.Println(<-leftmost)

}
