package main

import (
	"bufio"
	"fmt"
	"os"
)

type Stack struct {
	maxSize    int
	stackArray []byte
	top        int
}

func NewStack(s int) *Stack {
	return &Stack{
		maxSize:    s,
		stackArray: make([]byte, s),
		top:        -1,
	}
}

func (s *Stack) push(j byte) {
	s.top++
	s.stackArray[s.top] = j
}

func (s *Stack) pop() byte {
	value := s.stackArray[s.top]
	s.top--
	return value
}

func (s *Stack) peek() byte {
	return s.stackArray[s.top]
}

func (s *Stack) isEmpty() bool {
	return s.top == -1
}

func (s *Stack) isFull() bool {
	return s.top == s.maxSize-1
}

type Reverser struct {
	input  string
	output string
}

func NewReverser(in string) *Reverser {
	return &Reverser{
		input:  in,
		output: "",
	}
}

func (r *Reverser) doRev() string {
	stackSize := len(r.input)
	theStack := NewStack(stackSize)

	for j := 0; j < stackSize; j++ {
		ch := r.input[j]
		theStack.push(ch)
	}

	for !theStack.isEmpty() {
		ch := theStack.pop()
		r.output += string(ch)
	}

	return r.output
}

func getString() string {
	fmt.Print("Enter a string: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func main() {
	for {
		input := getString()
		if input == "" {
			break
		}

		theReverser := NewReverser(input)
		output := theReverser.doRev()
		fmt.Println(output)
	}
}