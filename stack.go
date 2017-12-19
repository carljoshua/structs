package structs

import (
    "reflect"
    "fmt"
)

type Stack struct {
    data    []interface{}
    dtype   reflect.Kind
    len     int
}

// NewStack() returns a Stack
func NewStack(dt reflect.Kind, l int) *Stack{
    return &Stack{dtype: dt, len: l}
}

// Push() adds an element into the Stack if the Stack is not full
// It also checks the data type of the value before storing it in the stack
func (s *Stack) Push(v interface{}){
    if !s.IsFull(){
        vt := reflect.TypeOf(v)
        if (vt.Kind() == s.dtype){
            s.data = append(s.data, v)
        }
    }
}

// Pop() removes the top-most element in the Stack
func (s *Stack) Pop(){
    if !s.IsEmpty(){
        s.data = s.data[:len(s.data) - 1]
    }
}

// Top() return the top-most element in the Stack
func (s *Stack) Top() interface{} {
    if !s.IsEmpty(){
        return s.data[len(s.data) - 1]
    }
    return nil
}

// IsEmpty() returns true if the Stack is empty, otherwise false
func (s *Stack) IsEmpty() bool {
    if len(s.data) == 0{
        return true
    }
    return false
}

// IsFull() returns true if the Stack is full, otherwise false
func (s *Stack) IsFull() bool {
    if len(s.data) >= s.len{
        return true
    }
    return false
}

// Print() prints the data in the stack into the console
func (s *Stack) Print() {
    fmt.Println(s.data)
}
