package structs

import (
    "reflect"
    "fmt"
)

// Linked List Head
type Head struct {
    next        *ListNode
    len         int
    dtype       reflect.Kind
}

// Linked List Node
type ListNode struct {
    data        interface{}
    next        *ListNode
}

// NewList() creates the head node of the Linked List
func NewList(d_type reflect.Kind) *Head{
    return &Head{len: 0, dtype: d_type}
}

// Insert() adds the an element into the given index
// If the index exceeds the length of the list,
// the element will be added at the last
func (l *Head) Insert(key int, v interface{}) {
    if l.next == nil {
        new_node := &ListNode{data: v}
        l.next = new_node
    }else if key == 0{
        new_node := &ListNode{data: v, next: l.next}
        l.next = new_node
    }else {
        cur_node := l.next
        for i := 0; i < key - 1 && cur_node.next != nil; i++ {
            cur_node = cur_node.next
        }
        new_node := &ListNode{data: v, next: cur_node.next}
        cur_node.next = new_node
    }
    l.len += 1
}

// Remove() removes an element in the given index
func (l *Head) Remove(key int) {
    if key + 1 > l.len{

    }else if key == 0 {
        temp := l.next
        l.next = temp.next
    }else{
        tmp := l.next
        for i := 0; i < key - 1; i++{
            tmp = tmp.next
        }
        tmp2 := tmp.next
        tmp.next = tmp2.next
    }
}

// Get() returns the element in the given index
func (l *Head) Get(key int) interface{}{
    if key > l.len - 1{
        //sample error
        panic("index out of bounds")
    }

    temp := l.next
    for i := 0; i < key; i++{
        temp = temp.next
    }
    return temp.data
}

// Print() prints the data in the list into the console
func (l *Head) Print() {
    cur_node := l.next
    i := false
    fmt.Print("[")
    for cur_node != nil {
        if i {
            fmt.Print(" ")
        }
        if !i {
            i = true
        }
        fmt.Print(cur_node.data)
        cur_node = cur_node.next
    }
    fmt.Print("]\n")
}
