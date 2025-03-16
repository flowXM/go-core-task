package main

import (
	"reflect"
	"testing"
	"time"
)

func TestPipe(t *testing.T) {
	input := make(chan uint8)
	output := make(chan float64)

	go pipe(input, output)

	actual := make([]float64, 0)
	expected := []float64{4, 36, 9}

	go func() {
		for num := range output {
			actual = append(actual, num)
		}
	}()

	input <- 2
	input <- 6
	input <- 3

	time.Sleep(time.Second * 5)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Incorrect pipe")
	}
}
