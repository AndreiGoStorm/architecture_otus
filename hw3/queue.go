package hw3

import "container/list"

type Queue struct {
	list *list.List
}

func NewQueue() *Queue {
	return &Queue{list.New()}
}

func (q *Queue) Enqueue(e interface{}) {
	q.list.PushBack(e)
}

func (q *Queue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}

	e := q.list.Front()
	q.list.Remove(e)
	return e.Value
}

func (q *Queue) Len() int {
	return q.list.Len()
}

func (q *Queue) IsEmpty() bool {
	return q.list.Len() == 0
}
