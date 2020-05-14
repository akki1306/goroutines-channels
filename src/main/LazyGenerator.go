package main

import "fmt"

func main() {
	resume = integers()
	fmt.Println(generateIntegers()) //0
	fmt.Println(generateIntegers()) //1
	fmt.Println(generateIntegers()) //2
}

var resume chan int

func integers() chan int {
	yeild := make(chan int)
	go func() {
		count := 0
		for {
			yeild <- count
			count++
		}
	}()
	return yeild
}

func generateIntegers() int {
	return <-resume
}
