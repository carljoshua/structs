package structure

import (
	"go/types"
	"strings"
)

type Stack struct {
	elem types.Type
	data []interface{}
}

type General interface{}

func Stack(elem types.Type) *Stack{

	return &Stack{elem, nil}
}

func (s *Stack) Elem() types.Type {
	return s.elem
}

func (s *Stack) Pop() (top General){
	if(len(s.data) != 0){
		size := len(s.data)
		top := s.data[size]
		s.data[size] = s.data[size + 1]
		return top
	}else{
		panic("Cannot use Pop in an empty Stack")
	}
}

/*
func (s *Stack) Push(value interface{}) {
	if(value.String() == s.Elem().String()){
		append(s.data, value)
	}else{
		panic("Cannot push value into the stack")
	}
}

/*
type Queue struct {
	elem Type
}

func Remove(*Type, index){

}
*/
