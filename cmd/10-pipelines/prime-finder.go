package main

func PrimeFinder(done <-chan bool, integerStream <-chan int) <-chan int {
	primeStream := make(chan int)

	go func() {
		defer close(primeStream)

		for {
			select {
			case <-done:
				return
			case value := <-integerStream:
				if IsPrime(value) {
					primeStream <- value
				}
			}
		}
	}()

	return primeStream
}

func IsPrime(number int) bool {
	for i := number - 1; i > 1; i-- {
		if number%i == 0 {
			return false
		}
	}
	return true
}
