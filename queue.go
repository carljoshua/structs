package structs

import "fmt"

// Structure of the Queue
type Queue struct {
    data []interface{}
    size int
}

// NewQueue creates a Queue
func NewQueue(l int) *Queue{
    return &Queue{size: l}
}

// Enqueue adds an element inside the Queue if the Queue is not full
func (q *Queue) Enqueue(v interface{}){
    if !q.IsFull(){
        q.data = append(q.data, v)
    }
}

// Dequeue removes the bottom-most element in the Queue
func (q *Queue) Dequeue(){
    if !q.IsEmpty(){
        q.data = q.data[1:]
    }
}

// Peek returns the bottom-most element in the Queue
func (q *Queue) Peek() interface{}{
    if !q.IsEmpty(){
        return q.data[0]
    }
    return nil
}

// Print prints the elements inside the queue using fmt
func (q *Queue) Print(){
    fmt.Print("[ ")
    for _, val := range q.data{
        fmt.Print(val)
        fmt.Print(" ")
    }
    fmt.Print("]\n")
}

// IsEmpty returns true if the Queue is empty, otherwise false
func (q *Queue) IsEmpty() bool{
    if len(q.data) == 0{
        return true
    }
    return false
}

// IsFull returns true if the Queue is full, otherwise false
func (q *Queue) IsFull() bool{
    if len(q.data) >= q.size{
        return true
    }
    return false
}
