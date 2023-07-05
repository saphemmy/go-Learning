# Go-Learning


## Goroutines Vs Threads

- Go doesn't use threads, but it's important to understand CPU threads to gen an idea of what concurrency is, especially in the other programming languages

- A Goroutine is a function or method whic executes independently and simultaneously in connect with any other Goroutines in your program. Or in other words, every concurrently executing activity in Go language is know as Goroutines.

- A thread is a lightweight process, or in other words, a thread is a unit which exceutes the code under the program. So every program has logis and a thread is responsible for executing this logic.

- Goroutines are lighter than OS threads

- It's up to Go to the "heavy lifting" with Goroutines

- Goroutines run... Multiple Goroutines on ONE thread

- Goroutines have faster startup times.

- Goroutines are concurrency but just a different type of concurrency. "threading 2.0", if you will.

## Channels

- Channels are like pipe. One Goroutines puts data ON TO the channel, and another Goroutine grabs it off.

- Two types of channels: Buffered and Unbuffered.


### Buffed and Unbuffered

### Buffered
- Can hold data
- A goroutine put data on a buffered channel, then move onto a new task

### Unbuffered
- Cannot hold data
- Any Goroutine that puts data on a channel is blocked until there's a receiver on the other end, forcing synchronous behavior