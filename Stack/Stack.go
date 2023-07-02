package main

import "fmt"

type Stack struct {
	maxSize    int      // Stack size
	stackArray []int64
	top        int      // Top of the stack
}

// Constructor
func NewStack(s int) *Stack {
	return &Stack{
		maxSize:    s,                //  Determining the stack size
		stackArray: make([]int64, s), // Creating an array
		top:        -1,               // No elements by default.
	}
}

// Placing an element on top of the stack
func (s *Stack) Push(j int64) {
    // increase top and insert element
	s.top++
	s.stackArray[s.top] = j
}

// Popping an array from the top of the stack
func (s *Stack) Pop() int64 {
    // Removing an element and reducing top
	value := s.stackArray[s.top]
	s.top--
	return value
}

func (s *Stack) Peek() int64 { // Reading an element from the stack top
	return s.stackArray[s.top]
}

func (s *Stack) IsEmpty() bool {
	return s.top == -1 // true if the stack is empty
}

func (s *Stack) IsFull() bool {
	return s.top == s.maxSize-1 // true if the stack is full
}

func main() {
	theStack := NewStack(10) // Create a new stack
	theStack.Push(20)        // pushing elements on the stack
	theStack.Push(30)
	theStack.Push(40)
	theStack.Push(50)

    // Until the stack is empty
	for !theStack.IsEmpty() {
        value := theStack.Pop()
		fmt.Println(value)
	}
}