package stack

import (
	"iter"

	"github.com/X-xiexianhui/go-stl/list"
)

type Stack[T any] struct {
	top *list.List[T]
}

func New[T any]() *Stack[T] {
	return &Stack[T]{list.New[T]()}
}
func (s *Stack[T]) IsEmpty() bool {
	return s.top.Len() == 0
}
func (s *Stack[T]) Len() int {
	return s.top.Len()
}
func (s *Stack[T]) Push(val T) {
	s.top.PushFront(val)
}
func (s *Stack[T]) Pop() (val T) {
	if s.top.Len() == 0 {
		return
	}
	backNode := s.top.Front()
	val = backNode.Value
	s.top.Remove(backNode)
	return
}
func (s *Stack[T]) Top() (val T) {
	if s.top.Len() == 0 {
		return
	}
	backNode := s.top.Front()
	val = backNode.Value
	return
}
func (s *Stack[T]) Iterator() iter.Seq[T] {
	return s.top.Iterator()
}
