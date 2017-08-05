package structs

type ListNode struct {
    val interface{}
    next *ListNode
}

func NewList() *ListNode{
    return &ListNode{}
}

func (l *ListNode) Insert(v interface{}, key int){
    tmp := &ListNode{val: v}
    if key == 0{
        tmp.next = l
        l = tmp
        return
    }
    tmp2 := l
    for i := 0; i < key-1; i++{
        tmp2 = tmp2.next
    }
    tmp.next = tmp2.next
    tmp2.next = tmp
}

func (l *ListNode) Search(key int) interface{}{
    temp := l
    for i := 0; i < key; i++{
        temp = temp.next
    }
    return temp.val
}
