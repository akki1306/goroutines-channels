package main

import "fmt"

func generatorInt() chan int {
	out := make(chan int)
	go func() {
		for i := 0; i <= 10; i++ {
			out <- i
		}
		close(out)
	}()

	return out
}

func squareNumber(in chan int) <-chan int {
	out := make(chan int)
	go func() {
		for number := range in {
			out <- number * number
		}
		close(out)
	}()
	return out
}

func main() {
	ch := squareNumber(generatorInt())

	for ch != nil {
		if val, ok := <-ch; ok {
			fmt.Println(val)
		} else {
			ch = nil
		}
	}
}
