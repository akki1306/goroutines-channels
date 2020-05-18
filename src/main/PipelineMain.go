package main

import (
	"math"

	"pipeline"
)

func generateNumber() chan interface{} {
	out := make(chan interface{})
	go func() {
		defer close(out)
		for i := 1; i <= 10; i++ {
			out <- i
		}
	}()
	return out
}

func squareNumber(number interface{}) (interface{}, error) {
	var err error
	/*	if number == 3 {
			err = errors.New("this is a test error")
		}
	*/return number.(int) * number.(int), err
}

func squareRootNumber(number interface{}) (interface{}, error) {
	var err error
	/*	if number == 25 {
		err = errors.New("this is a test error")
	}*/
	return math.Sqrt(float64(number.(int))), err
}

func main() {
	pipeline := pipeline.Pipeline{ErrorChan: make(chan interface{}), Processors: []pipeline.ProcessorFunc{squareNumber, squareRootNumber}}
	pipeline.InitPipeline(generateNumber())
	pipeline.PrintResults()
}
