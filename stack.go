package structs

import "fmt"

// Structure of the Stack
type Stack struct {
    data []interface{}
    size int
}

// NewStack creates a Stack
func NewStack(l int) *Stack{
    return &Stack{size: l}
}

// Push adds an element into the Stack if the Stack is not full
func (s *Stack) Push(v interface{}){
    if !s.IsFull(){
        s.data = append(s.data, v)
    }
}

// Pop removes the top-most element in the Stack
func (s *Stack) Pop(){
    if !s.IsEmpty(){
        s.data = s.data[:len(s.data) - 1]
    }
}

// Top return the top-most element in the Stack
func (s *Stack) Top() interface{}{
    if !s.IsEmpty(){
        return s.data[len(s.data) - 1]
    }
    return nil
}

// Print prints the data inside the Stack using fmt package
func (s *Stack) Print(){
    fmt.Print("[ ")
    for _, val := range s.data{
        fmt.Print(val)
        fmt.Print(" ")
    }
    fmt.Print("]\n")
}

// IsEmpty returns true if the Stack is empty, otherwise false
func (s *Stack) IsEmpty() bool{
    if len(s.data) == 0{
        return true
    }
    return false
}

// IsFull returns true if the Stack is full, otherwise false
func (s *Stack) IsFull() bool{
    if len(s.data) >= s.size{
        return true
    }
    return false
}
