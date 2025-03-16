package main

import (
	"reflect"
	"slices"
	"testing"
)

func TestMergeChannels(t *testing.T) {
	chan1 := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			chan1 <- i
		}
		close(chan1)
	}()

	chan2 := make(chan int)
	go func() {
		for i := 5; i < 10; i++ {
			chan2 <- i
		}
		close(chan2)
	}()

	chan3 := make(chan int)
	go func() {
		for i := 10; i < 15; i++ {
			chan3 <- i
		}
		close(chan3)
	}()

	chan4 := make(chan int)
	go func() {
		for i := 15; i < 20; i++ {
			chan3 <- i
		}
		close(chan4)
	}()

	merged := mergeChannels(chan1, chan2, chan3, chan4)

	var actual []int
	for i := range merged {
		actual = append(actual, i)
	}
	slices.Sort(actual)

	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Incorrect channel merge got: %v, expected: %v", actual, expected)
	}
}
