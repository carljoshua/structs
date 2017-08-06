package structs

import "fmt"

type ListNode struct {
    val interface{}
    next *ListNode
    isEmpty bool
}

func NewList() *ListNode{
    return &ListNode{isEmpty: true}
}

func (l *ListNode) Insert(v interface{}, key int){
    if l.isEmpty{
        l.val = v
        l.isEmpty = false
    }else{
        tmp := &ListNode{val: v}
        if key == 0{
            b_val := l.val
            b_next := l.next

            l.val = tmp.val
            tmp.val = b_val
            tmp.next = b_next
            l.next = tmp
        }else{
            tmp2 := l
            for i := 0; i < key - 1; i++{
                tmp2 = tmp2.next
            }
            tmp.next = tmp2.next
            tmp2.next = tmp
        }
    }
}

func (l *ListNode) Remove(key int) {
    if key == 0{
        tmp := l.next
        l.val = tmp.val
        l.next = tmp.next
    }else{
        tmp := l
        for i := 0; i < key - 1; i++{
            tmp = tmp.next
        }
        tmp2 := tmp.next
        tmp.next = tmp2.next
    }
}

func (l *ListNode) Search(key int) interface{}{
    temp := l
    for i := 0; i < key; i++{
        temp = temp.next
    }
    return temp.val
}

func (l *ListNode) Print(){
    temp := l
    for temp != nil{
        fmt.Printf("%s\n", temp)
        temp = temp.next
    }
}
