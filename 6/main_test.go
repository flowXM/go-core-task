package main

import (
	"testing"
)

func TestRandGenerator(t *testing.T) {
	randChan := make(chan int)
	minVal, maxVal := 0, 100000
	go randGenerator(randChan, minVal, maxVal)

	val1 := <-randChan

	if val1 < minVal || val1 > maxVal {
		t.Errorf("Invalid value: %d", val1)
	}
	val2 := <-randChan

	if val2 < minVal || val2 > maxVal {
		t.Errorf("Invalid value: %d", val2)
	}

	if val1 == val2 {
		t.Errorf("Invalid random generator or probability hit: %g%%", (1.0/float64(maxVal-minVal))*100)
	}
}
