package main

import (
	"fmt"
	"math"
	"time"
)

type WaitGroup struct {
	ch     chan struct{}
	waitCh []chan struct{}
}

func NewWaitGroup() WaitGroup {
	return WaitGroup{ch: make(chan struct{}, math.MaxInt), waitCh: make([]chan struct{}, 0)}
}

func (w *WaitGroup) Add(delta int) {
	for i := 0; i < delta; i++ {
		w.ch <- struct{}{}
	}
}

func (w *WaitGroup) Done() {
	<-w.ch
	if len(w.ch) == 0 {
		for _, wch := range w.waitCh {
			<-wch
		}
	}
}

func (w *WaitGroup) Wait() {
	wc := make(chan struct{})
	w.waitCh = append(w.waitCh, wc)
	wc <- struct{}{}
}

func main() {
	wg := NewWaitGroup()

	solve := func(delay time.Duration) {
		wg.Add(1)
		time.Sleep(time.Second * delay)
		wg.Done()
		fmt.Println("func done")
	}

	go solve(2)
	go solve(4)
	go solve(6)
	go solve(7)
	go solve(9)

	wg.Wait()

	fmt.Println("All done")
}
