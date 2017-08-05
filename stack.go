package structs

type Stack struct {
    data []interface{}
}

func NewStack() *Stack{
    return &Stack{}
}

func (s *Stack) Push(v interface{}){
    s.data = append(s.data, v)
}

func (s *Stack) Pop(){
    if !s.IsEmpty(){
        s.data = s.data[:len(s.data) - 1]
    }
}

func (s *Stack) Top() interface{}{
    if !s.IsEmpty(){
        return s.data[len(s.data) - 1]
    }
    return nil
}

func (s *Stack) IsEmpty() bool{
    if len(s.data) == 0{
        return true
    }
    return false
}
