// Package snickerboa Stack of []interface{} methods: NewStack, Push, Pop, Peek, IsEmpty, Size
// Source chat.openai.com
package snickerboa

import (
	"errors"
)

// Stack is a generic stack data structure.
type Stack struct {
	items []interface{}
}

// NewStack creates a new stack.
func NewStack() *Stack {
	return &Stack{
		items: []interface{}{},
	}
}

// Push adds an item to the top of the stack.
func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top item from the stack.
// If the stack is empty, it returns an error.
func (s *Stack) Pop() (interface{}, error) {
	if len(s.items) == 0 {
		return nil, errors.New("stack is empty")
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, nil
}

// Peek returns the top item from the stack without removing it.
// If the stack is empty, it returns an error.
func (s *Stack) Peek() (interface{}, error) {
	if len(s.items) == 0 {
		return nil, errors.New("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

// IsEmpty returns true if the stack is empty, and false otherwise.
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of items in the stack.
func (s *Stack) Size() int {
	return len(s.items)
}

/*
 * Usage Exampel
 *
func main() {
	s := NewStack()
	fmt.Println(s.IsEmpty()) // true
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println(s.Size()) // 3
	item, err := s.Peek()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(item) // 3
	}
	item, err = s.Pop()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(item) // 3
	}
	fmt.Println(s.Size())    // 2
	fmt.Println(s.IsEmpty()) // false
}

*/
