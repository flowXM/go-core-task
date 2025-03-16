package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func randGenerator(randChan chan int, min, max int) <-chan int {
	for {
		rndNum := rand.IntN(max-min) + min
		randChan <- rndNum
	}
}

func main() {
	randChan := make(chan int)
	go randGenerator(randChan, 0, 500)
	for {
		select {
		case res := <-randChan:
			fmt.Println(res)
		}
		time.Sleep(time.Second)
	}
}
