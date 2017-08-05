package structs

type Node struct {
    val int
    l_child *Node
    r_child *Node
    isEmpty bool
}

func NewTree() *Node{
    return &Node{isEmpty: true}
}

func (n *Node) Insert(v int) *Node{
    if n == nil{
        n = &Node{val: v}
        return n
    }else if n.isEmpty{
        n.val = v
        n.isEmpty = false
    }else if v < n.val{
        n.l_child = n.l_child.Insert(v)
    }else if v > n.val{
        n.r_child = n.r_child.Insert(v)
    }
    return n
}

func (n *Node) Remove(){
    
}
