package structs

import (
    "reflect"
    "fmt"
)

type Queue struct {
    data    []interface{}
    dtype   reflect.Kind
    len     int
}

// NewQueue() creates a Queue
func NewQueue(dt reflect.Kind, l int) *Queue{
    return &Queue{dtype: dt, len: l}
}

// Enqueue() adds an element inside the Queue if the Queue is not full
func (q *Queue) Enqueue(v interface{}){
    if !q.IsFull(){
        vt := reflect.TypeOf(v)
        if (vt.Kind() == q.dtype){
            q.data = append(q.data, v)
        }
    }
}

// Dequeue() removes the bottom-most element in the Queue
func (q *Queue) Dequeue(){
    if !q.IsEmpty(){
        q.data = q.data[1:]
    }
}

// Peek() returns the bottom-most element in the Queue
func (q *Queue) Peek() interface{}{
    if !q.IsEmpty(){
        return q.data[0]
    }
    return nil
}

// IsEmpty() returns true if the Queue is empty, otherwise false
func (q *Queue) IsEmpty() bool{
    if len(q.data) == 0{
        return true
    }
    return false
}

// IsFull() returns true if the Queue is full, otherwise false
func (q *Queue) IsFull() bool{
    if len(q.data) >= q.len{
        return true
    }
    return false
}

// Print() prints the data in the queue into the console
func (q *Queue) Print() {
    fmt.Println(q.data)
}
