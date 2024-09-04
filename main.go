package main

import (
	datastructures "github.com/johnifegwu/Go-Generic-DataStructures/doublyLinkedList"
)

func main() {
	// This is how your code will be called.
	// Your answer should respect the linked list order.
	// You can edit this code to try different testing cases.
	testCases := []struct {
		index int
		value string
	}{
		{index: 0, value: "C"},
		{index: 0, value: "A"},
		{index: 1, value: "B"},
		{index: 3, value: "D"},
	}
	dl := &datastructures.DoublyLinkedList[string]{}
	err := dl.AddElements(testCases)
	forwardPrint := dl.PrintForward()
	reversePrint := dl.PrintReverse()
}
