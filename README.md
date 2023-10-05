# go-concurrency
some research on go concurrency

## The Go approach
Don't communicate by sharing memory, share memory by communicating.

## The Two Principles
- Start goroutines when you have concurrent work.
- Share by communicating.


## Channel

Go 的开发者极力推荐使用 Channel，不过，这两年，大家意识到，Channel 并不是处理并发问题的“银弹”，有时候使用并发原语更简单，而且不容易出错。所以，我给你提供一套选择的方法:
- 共享资源的并发访问使用传统并发原语；
- 复杂的任务编排和消息传递使用 Channel；
- 消息通知机制使用 Channel，除非只想 signal 一个 goroutine，才使用 Cond；
- 简单等待所有任务的完成用 WaitGroup，也有 Channel 的推崇者用 Channel，都可以；
- 需要和 Select 语句结合，使用 Channel；
- 需要和超时配合时，使用 Channel 和 Context。

## Reference
- [Go Concurrency Patterns](https://talks.golang.org/2012/concurrency.slide#1)
and [source](https://talks.golang.org/2012/concurrency/support/)
- [Advanced Go Concurrency Patterns](https://talks.golang.org/2013/advconc.slide)