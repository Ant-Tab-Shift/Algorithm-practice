//go:build concurrency

package main

import (
	"fmt"
	rand "math/rand/v2"
)

type Runnable interface {
	Run() error
}

type RandomError struct {
	probability float64
}

func (re *RandomError) Run() error {
	if rand.Float64() < re.probability {
		return fmt.Errorf("congratulations! it's error, probability was %.3f", re.probability)
	}
	return nil
}

func NewRandomErorr(probability float64) *RandomError {
	return &RandomError{probability}
}

func main() {
	runners := make([]Runnable, 10)
	for i := range len(runners) {
		runners[i] = NewRandomErorr(rand.Float64())
	}

	answers := make([]error, len(runners))
	isReadyChanel := make(chan struct{}, len(answers))
	for i := range len(answers) {
		go func(idx int) {
			answers[idx] = runners[idx].Run()
			isReadyChanel <- struct{}{}
		}(i)
	}

	for range len(answers) {
		<-isReadyChanel
	}
	for i := range len(answers) {
		fmt.Println(answers[i])
	}
}
