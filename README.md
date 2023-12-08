# master-go-concurrency-patterns

Master Go Concurrency Patterns

## Concurrency Primitives

### Go Routines

![Fork-Join model](./assets/go-fork-join.png 'Fork-Join model')

**_Processes running concurrently on a separate thread from main._**

A goroutine is a lightweight thread managed by the Go runtime.

`go f(x, y, z)`

starts a new goroutine running

`f(x, y, z)`

The evaluation of f, x, y, and z happens in the current goroutine and the execution of f happens in the new goroutine.

Goroutines run in the same address space, so access to shared memory must be synchronized. The sync package provides useful primitives, although you won't need them much in Go as there are other primitives.

### Channels

**_Channels are FIFO queues typically used to communicate typed information between go routines._**

Channels are a typed conduit through which you can send and receive values with the channel operator, <-.

```
  ch <- v     // Send v to channel ch
  v := <- ch  // Receive from ch, and assign value to v
```

(The data flows in the direction of the arrow.)

Like maps and slices, channels must be created before use:

`ch := make(chan int)`

By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables.

The example code sums the numbers in a slice, distributing the work between two goroutines. Once both goroutines have completed their computation, it calculates the final result.

### Select

**_The select statement allows selecting the first operation ready to be executed from multiple channels. It can be a signal to perform some useful task or to exit the infinite loop._**

```
	select {
	case msgFromChannelOne := <-channelOne:
		fmt.Println(msgFromChannelOne)
	case msgFromChannelTwo := <-channelTwo:
		fmt.Println(msgFromChannelTwo)
	}
```

Select statement will block the execution until it receive data from a channel.

### WaitGroups

**_To wait for goroutines to finish, we can use wait groups._**

```
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			doSomethingConcurrently()
		}()
	}

	wg.Wait()
```

### Generators

A generator in computer science is a function, sometimes called a subroutine that produces values to iterate over. Those values can be the result of a computation. The iteration ends either when the generator stops producing values or when the caller terminates it explicitly (with the break keyword for instance).

## Concurrency Patterns

### for-select-done Pattern

The main idea of the for-select-done pattern is to use an infinite _for_ loop to handle events from various channels using the _select_ statement.

The _select_ statement allows selecting the first operation ready to be executed from multiple channels. It can be a signal to perform some useful task or to exit the infinite loop.

In this pattern, the infinite _for_ loop is usually called in a separate goroutine to avoid blocking the main thread.

### Pipeline Pattern

![Pipeline pattern example](./assets/pipeline-pattern-example.png 'Pipeline pattern example')

The pipeline pattern allows you to easily compose sequential stages by chaining stages.

### Fan-In Fan-Out Pattern

![Fan-In Fan-Out pattern example](./assets/find-primes-pipeline.png 'Fan-In Fan-Out pattern example')

The Fan-In Fan-Out pattern allows you to execute several instances of a stage in a pipeline.

## Links

- [Master Go Programming With These Concurrency Patterns (Video - part 1)](https://www.youtube.com/watch?v=qyM8Pi1KiiM)
- [Master Go Programming With These Concurrency Patterns (Video - part 2)](https://www.youtube.com/watch?v=wELNUHb3kuA)
- [Golang Concurrency Patterns](https://medium.com/@ninucium/golang-concurrency-patterns-for-select-done-errgroup-and-worker-pool-645bec0bd3c9)
- [Tour of Go - Concurrency](https://go.dev/tour/concurrency/1)
- [Intro to Concurrency in Golang](https://vietmle.com/posts/go_conc_intro/)
- [Golang closures and goroutines](https://go.dev/doc/faq#closures_and_goroutines)
- [Concurrent data pipelines in Golang](https://towardsdatascience.com/concurrent-data-pipelines-in-golang-85b18c2eecc2)
- [Generators With Go Channels](https://www.mickaelvieira.com/blog/2020/02/21/a-short-explanation-of-generators-with-go-channels.html)
- [Go Fan-In / Fan-Out pattern example](https://austburn.me/blog/a-better-fan-in-fan-out-example.html)
