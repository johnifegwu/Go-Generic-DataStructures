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
	dll := &datastructures.DoublyLinkedList[string]{}
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
