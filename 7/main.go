package main

import (
	"fmt"
	"sync"
)

func redirectChan(input <-chan int, output chan int, wg *sync.WaitGroup) {
	for message := range input {
		output <- message
	}

	wg.Done()
}

func mergeChannels(chans ...<-chan int) <-chan int {
	merged := make(chan int)

	wg := sync.WaitGroup{}
	wg.Add(len(chans))

	for _, ch := range chans {
		go redirectChan(ch, merged, &wg)
	}

	go func() {
		wg.Wait()
		close(merged)
	}()

	return merged
}

func main() {
	chan1 := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			chan1 <- i
		}
		close(chan1)
	}()

	chan2 := make(chan int)
	go func() {
		for i := 10; i < 20; i++ {
			chan2 <- i
		}
		close(chan2)
	}()

	merged := mergeChannels(chan1, chan2)

	for i := range merged {
		fmt.Println(i)
	}
}
