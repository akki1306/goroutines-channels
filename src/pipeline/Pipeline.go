package pipeline

import "fmt"

type ProcessorFunc func(in interface{}) (interface{}, error)

type Pipeline struct {
	ResultChan chan interface{}
	Processors []ProcessorFunc
	InputChan  chan interface{}
	ErrorChan  chan interface{}
}

func (pipeline *Pipeline) InitPipeline(in chan interface{}) {
	for i := 0; i < len(pipeline.Processors); i++ {
		in = pipeline.addPipelineMember(in, pipeline.Processors[i])
	}
	pipeline.ResultChan = in
}

func (pipeline *Pipeline) addPipelineMember(in chan interface{}, proFunc ProcessorFunc) chan interface{} {
	out := make(chan interface{})
	go func() {
		defer close(out)
		for in != nil {
			select {
			case number, ok := <-in:
				if ok {
					i, err := proFunc(number)
					if err != nil {
						pipeline.ErrorChan <- err
					}
					out <- i
				} else {
					in = nil
				}
			case err := <-pipeline.ErrorChan:
				pipeline.ErrorChan <- err
				return
			}
		}
	}()

	return out
}

func (pipeline *Pipeline) PrintResults() {
	for pipeline.ResultChan != nil {
		select {
		case number, ok := <-pipeline.ResultChan:
			if ok {
				fmt.Println("Number is ", number)
			} else {
				pipeline.ResultChan = nil
			}
		case err := <-pipeline.ErrorChan:
			fmt.Println("Error occurred while processing", err)
			close(pipeline.ResultChan)
			return
		}
	}
}
