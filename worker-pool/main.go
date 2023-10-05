package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

// 需要的时候再创建goroutine，避免不必要的空跑协程
// 使用 semaphore-channel 来控制创建的协程总数

type token struct{}

func runWork(tasks []string, limit int) <-chan bool {
	done := make(chan bool)
	go func() {
		sem := make(chan token, limit)
		for _, task := range tasks {
			sem <- token{}
			go func(task string) {
				time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
				fmt.Printf("I am doing taks: %s\n", task)
				<-sem
			}(task)
		}

		for n := limit; n > 0; n-- {
			sem <- token{}
		}

		done <- true
	}()

	return done
}

func main() {

	totalJob := 1000
	tasks := make([]string, totalJob)

	for i := 0; i < totalJob; i++ {
		tasks[i] = fmt.Sprintf("[task]-%d", i)
	}

	limit := 40

	done := runWork(tasks, limit)
	t := time.Tick(time.Duration(1) * time.Second)
	for {
		select {
		case <-t:
			num := runtime.NumGoroutine()
			fmt.Printf("---------------------------------total goroutine: %d\n", num)
		case <-done:
			goto jobDone
		}
	}

jobDone:
	fmt.Println("job done")

}
