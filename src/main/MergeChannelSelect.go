package main

import (
	"fmt"
	"time"
)

func main() {
	chEven := generateEven()
	chOdd := generateOdd()
	mergeCh := make(chan int)
	done := make(chan bool)

	defer close(chEven)
	defer close(chOdd)
	defer close(done)
	defer close(mergeCh)

	go func() {
		for {
			select {
			case number := <-chEven:
				mergeCh <- number
			case number := <-chOdd:
				mergeCh <- number
			case <-time.After(1000 * time.Millisecond):
				done <- true
				return
			}
		}
	}()

	go func() {
		for {
			select {
			case number := <-mergeCh:
				fmt.Println("Received Number :", number)
			case <-done:
				return
			}
		}
	}()

	time.Sleep(5000 * time.Millisecond)
}

func generateEven() chan int {
	ch := make(chan int)

	go func() {
		for i := 1; i < 20; i++ {
			if i%2 == 0 {
				ch <- i
			}
		}
	}()

	return ch
}

func generateOdd() chan int {
	ch := make(chan int)

	go func() {
		for i := 1; i < 20; i++ {
			if i%2 != 0 {
				ch <- i
			}
		}
	}()

	return ch
}
