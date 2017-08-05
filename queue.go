package structs

type Queue struct {
    data []interface{}
}

func NewQueue() *Queue{
    return &Queue{}
}

func (q *Queue) Push(v interface{}){
    q.data = append(q.data, v)
}

func (q *Queue) Pop(){
    if !q.IsEmpty(){
        q.data = q.data[1:]
    }
}

func (q *Queue) Peek() interface{}{
    if !q.IsEmpty(){
        return q.data[0]
    }
    return nil
}

func (q *Queue) IsEmpty() bool{
    if len(q.data) == 0{
        return true
    }
    return false
}
