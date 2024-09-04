package doublyLinkedList

// Write your answer here, and then test your code.
// Your job is to implement the Add(index int, v T) method.

import (
	"errors"
	"fmt"
)

// Change these boolean values to control whether you see
// the expected answer and/or hints.
const showExpectedResult = false
const showHints = false

var maxSize = 99

type Node[T any] struct {
	value      T
	next, prev *Node[T]
}

type DoublyLinkedList[T any] struct {
	size int
	head *Node[T]
	tail *Node[T]
}

// Add appends the given value at the given index.
// Returns an error in the case that the index exceeds the list size.
func (dll *DoublyLinkedList[T]) Add(index int, v T) error {

	//Check if max size is reached
	if index < 0 || index > maxSize {
		return errors.New("index exceeds list size")
	}

	//Add to list
	newNode := &Node[T]{value: v}

	if index == 0 {

		// Insert at the beginning
		if dll.head == nil {
			// Empty list
			dll.head = newNode
			dll.tail = newNode
		} else {
			newNode.next = dll.head
			dll.head.prev = newNode
			dll.head = newNode
		}

	} else if index == dll.size {
		// Insert at the end
		newNode.prev = dll.tail
		dll.tail.next = newNode
		dll.tail = newNode

	} else {
		// Insert in the middle
		current := dll.head
		for i := 0; i < index; i++ {
			current = current.next
		}
		newNode.next = current
		newNode.prev = current.prev
		current.prev.next = newNode
		current.prev = newNode
	}

	// increment size
	dll.size++
	return nil
}

func (l *DoublyLinkedList[T]) AddElements(elements []struct {
	index int
	value T
}) error {
	for _, e := range elements {
		if err := l.Add(e.index, e.value); err != nil {
			return err
		}
	}

	return nil
}

func (l *DoublyLinkedList[T]) PrintForward() string {
	if l.size == 0 {
		return ""
	}
	current := l.head
	output := "HEAD"
	for current != nil {
		output = fmt.Sprintf("%s -> %v", output, current.value)
		current = current.next
	}

	return fmt.Sprintf("%s -> NULL", output)
}

func (l *DoublyLinkedList[T]) PrintReverse() string {
	if l.size == 0 {
		return ""
	}
	current := l.tail
	output := "NULL"
	for current != nil {
		output = fmt.Sprintf("%s <- %v", output, current.value)
		current = current.prev
	}
	return fmt.Sprintf("%s <- HEAD", output)
}
