package main

import (
	"fmt"
	"math"
	"time"
)

func pipe(input <-chan uint8, output chan float64) {
	for num := range input {
		output <- math.Pow(float64(num), 2)
	}
}

func main() {
	input := make(chan uint8)
	output := make(chan float64)

	go pipe(input, output)

	go func() {
		for num := range output {
			fmt.Println(num)
		}
	}()

	input <- 2
	input <- 6
	input <- 3

	time.Sleep(time.Second * 5)
}
