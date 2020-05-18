package pipeline

import (
	"fmt"
)

type GeneratorFunc func() chan interface{}
type ProcessorFunc func(in chan interface{}) chan interface{}

type Pipeline struct {
	ResultChan chan interface{}
	Processors []ProcessorFunc
	Generator  GeneratorFunc
}

func (pipeline *Pipeline) BuildPipeline() {
	var ch chan interface{}
	ch = pipeline.Generator()
	for i := 0; i < len(pipeline.Processors); i++ {
		if i == 0 {
			ch = pipeline.Processors[i](ch)
		} else {
			ch = pipeline.Processors[i-1](ch)
		}
	}
	pipeline.ResultChan = ch
}

func (pipeline *Pipeline) PrintResults() {
	for number := range pipeline.ResultChan {
		fmt.Println(number)

	}

}
