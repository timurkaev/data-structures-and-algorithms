package main

import "fmt"

type Stack struct {
	maxSize    int
	stackArray []rune
	top        int
}

func NewStack(s int) *Stack {
	return &Stack{
		maxSize:    s,
		stackArray: make([]rune, s),
		top:        -1,
	}
}

func (s *Stack) push(j rune) {
	s.top++
	s.stackArray[s.top] = j
}

func (s *Stack) pop() rune {
	element := s.stackArray[s.top]
	s.top--
	return element
}

func (s *Stack) peek() rune {
	return s.stackArray[s.top]
}

func (s *Stack) isEmpty() bool {
	return s.top == -1
}

func checkParentheses(input string) bool {
	theStack := NewStack(len(input))

	for _, ch := range input {
		if ch == '(' || ch == '{' || ch == '[' {
			theStack.push(ch)
		} else if ch == ')' || ch == '}' || ch == ']' {
			if theStack.isEmpty() {
				return false
			}

			top := theStack.pop()
			if (ch == ')' && top != '(') ||
				(ch == '}' && top != '{') ||
				(ch == ']' && top != '[') {
				return false
			}
		}
	}

	return theStack.isEmpty()
}
func main() {
	input := "{(2+3)*[5-2]}"
	isBalanced := checkParentheses(input)

	if isBalanced {
		fmt.Println("Parentheses are balanced")
	} else {
		fmt.Println("Parentheses are not balanced")
	}
}