package rollingmean

// Write your answer here, and then test your code.
// Your job is to implement the RollingMean() method.

import (
	"fmt"
	"strings"
)

// Change these boolean values to control whether you see
// the expected answer and/or hints.
//const showExpectedResult = false
//const showHints = false

const WindowSize = 3

type NumericType interface {
	int64 | float64
}

type MovingAverage[T NumericType] struct {
	Position     int
	WindowValues [WindowSize]T
	WindowFilled bool
	Input        <-chan T
	Output       chan<- string
}

func NewMovingAverage[T NumericType](in <-chan T,
	out chan<- string) *MovingAverage[T] {
	return &MovingAverage[T]{
		Input:  in,
		Output: out,
	}
}

// CalculateMean prints the window and its calculated mean.
func (ma *MovingAverage[T]) CalculateMean() string {
	values := make([]string, WindowSize)
	var sum float64
	for i, v := range ma.WindowValues {
		values[i] = fmt.Sprint(v)
		sum += float64(v)
	}
	mean := sum / WindowSize
	return fmt.Sprintf("[%s] = %.2f", strings.Join(values, ","), mean)
}

func (ma *MovingAverage[T]) RollingMean() {
	for {
		val, ok := <-ma.Input
		if !ok {
			close(ma.Output)
			return
		}

		ma.WindowValues[ma.Position] = val
		ma.Position = (ma.Position + 1) % WindowSize

		// we have enough to cal our mean
		if ma.WindowFilled {
			ma.Output <- ma.CalculateMean()
		}

		// detect the first time we filled the window
		if !ma.WindowFilled {
			ma.WindowFilled = ma.Position == (WindowSize - 1)
		}
	}
}

func ReadResults(output <-chan string) []string {
	var results []string
	for r := range output {
		results = append(results, r)
	}
	return results
}
