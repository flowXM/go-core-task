package main

import (
	"sync"
	"testing"
	"time"
)

func TestNewWaitGroup(t *testing.T) {
	wg := NewWaitGroup()
	if wg.ch == nil {
		t.Errorf("Incorrect init")
	}
}

func TestWaitGroup_Add(t *testing.T) {
	wg := NewWaitGroup()
	wg.Add(1)
	if len(wg.ch) != 1 {
		t.Errorf("Incorrect add")
	}
}

func TestWaitGroup_Done(t *testing.T) {
	wg := NewWaitGroup()
	wg.Add(1)
	wg.Done()
	if len(wg.ch) != 0 {
		t.Errorf("Incorrect done")
	}
}

func TestWaitGroup_Wait(t *testing.T) {
	wg := NewWaitGroup()
	mx := sync.Mutex{}
	done := 0

	solve := func(delay time.Duration) {
		wg.Add(1)
		time.Sleep(time.Second * delay)
		wg.Done()
		mx.Lock()
		done++
		mx.Unlock()
	}

	go solve(2)
	go solve(4)
	go solve(6)
	go solve(7)
	go solve(9)

	wg.Wait()

	if done != 5 {
		t.Errorf("Incorrect wait")
	}
}
