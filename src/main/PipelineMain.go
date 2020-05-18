package main

import (
	"fmt"

	"pipeline"
)

func generatorInt() chan interface{} {
	out := make(chan interface{})
	go func() {
		defer close(out)
		for i := 0; i <= 10; i++ {
			out <- i
		}
	}()

	return out
}

func squareNumber(in chan interface{}) chan interface{} {
	out := make(chan interface{})
	go func() {
		defer close(out)
		for number := range in {
			out <- number.(int) * number.(int)
		}
	}()
	return out
}

func main() {
	end := make(chan interface{})
	pipeline := pipeline.Pipeline{ResultChan: end, Generator: generatorInt, Processors: []pipeline.ProcessorFunc{squareNumber}}
	pipeline.BuildPipeline()
	pipeline.PrintResults()

	fmt.Println("Successfully finished!!")
}
