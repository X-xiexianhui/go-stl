package queue

import (
	"iter"

	"github.com/X-xiexianhui/go-stl/list"
)

type Queue[T any] struct {
	queue *list.List[T]
}

func New[T any]() *Queue[T] {
	return &Queue[T]{list.New[T]()}
}
func (q *Queue[T]) IsEmpty() bool {
	return q.queue.Len() == 0
}
func (q *Queue[T]) Len() int {
	return q.queue.Len()
}
func (q *Queue[T]) Enqueue(val T) {
	q.queue.PushBack(val)
}
func (q *Queue[T]) Dequeue() (val T) {
	if q.queue.Len() == 0 {
		return
	}
	frontNode := q.queue.Front()
	val = frontNode.Value
	q.queue.Remove(frontNode)
	return
}
func (q *Queue[T]) Front() (val T) {
	if q.queue.Len() == 0 {
		return
	}
	frontNode := q.queue.Front()
	val = frontNode.Value
	return
}
func (q *Queue[T]) Rear() (val T) {
	if q.queue.Len() == 0 {
		return
	}
	backNode := q.queue.Back()
	val = backNode.Value
	return
}
func (q *Queue[T]) Iterator() iter.Seq[T] {
	return q.queue.Iterator()
}
