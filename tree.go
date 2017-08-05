package structs

type Node struct {
    value interface
    l_child *Node
    r_child *Node
}

type Tree struct {
    data []Node
}

func NewTree() *Tree{
    return *Tree{}
}

func (t *Tree) Insert(v interface){
    if t.IsEmpty(){
        root := &Node{value: v}
        t.data = append(t.data, root)
    }else{
        
    }
}

func (t *Tree) IsEmpty(){
    if len(t.data) == 0{
        return true
    }
    return false
}
