package main

import (
	datastructures "github.com/johnifegwu/Go-Generic-DataStructures/doublyLinkedList"
	dp "github.com/johnifegwu/Go-Generic-DataStructures/dp"
	rollingmean "github.com/johnifegwu/Go-Generic-DataStructures/rollingmean"
)

func main() {

	// This is how your code will be called.
	// Your answer should be the number of paths with the given cost.
	// You can edit this code to try different testing cases.
	cost := 8
	maze := dp.NewMaze([][]int{
		{1, 2, 1},
		{6, 1, 1},
		{4, 3, 3},
	})
	dp.CountPaths(maze, 0, 0, cost)

	// This is how your code will be called.
	// Your answer should be the sets of rolling means.
	numbers := []int64{5, 7, 9, 6, 8, 10}
	input := make(chan int64)
	output := make(chan string)
	movingAvg := rollingmean.NewMovingAverage(input, output)
	go movingAvg.RollingMean()
	go func() {
		for _, n := range numbers {
			input <- n
		}
		close(input)
	}()

	rollingmean.ReadResults(output)

	// This is how your code will be called.
	// Your answer should respect the linked list order.
	dll := &datastructures.DoublyLinkedList[string]{}

	testCases := []struct {
		index int
		value string
	}{
		{index: 0, value: "C"},
		{index: 0, value: "A"},
		{index: 1, value: "B"},
		{index: 3, value: "D"},
	}

	// Adding values to the list
	dll.AddElements(testCases)

	_ = dll.Add(0, "A") // Insert A at index 0
	_ = dll.Add(1, "B") // Insert B at index 1
	_ = dll.Add(1, "C") // Insert C at index 1
	_ = dll.Add(3, "D") // Insert D at index 3
	_ = dll.Add(0, "E") // Insert E at index 0

	dll.PrintForward()
	dll.PrintReverse()
	// Print the list to verify
	dll.PrintList()
}
