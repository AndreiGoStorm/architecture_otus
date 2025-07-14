package hw7

import (
	"container/list"
	"sync"
	"sync/atomic"
)

type Queue struct {
	mu    sync.Mutex
	cond  *sync.Cond
	list  *list.List
	count int32
	stop  bool
}

func NewQueue() *Queue {
	q := &Queue{list: list.New()}
	q.cond = sync.NewCond(&q.mu)
	return q
}

func (q *Queue) Enqueue(cmd Command) {
	q.mu.Lock()
	q.list.PushBack(cmd)
	q.mu.Unlock()
	atomic.AddInt32(&q.count, 1)

	q.cond.Signal()
}

func (q *Queue) Dequeue() Command {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.Wait()

	if !q.stop {
		cmd := q.list.Front()
		q.list.Remove(cmd)
		atomic.AddInt32(&q.count, -1)
		return cmd.Value.(Command)
	}
	return nil
}

func (q *Queue) Size() int {
	return int(atomic.LoadInt32(&q.count))
}

func (q *Queue) Wait() {
	for q.Size() == 0 && !q.stop {
		q.cond.Wait()
	}
}

func (q *Queue) Stop() {
	q.mu.Lock()
	q.stop = true
	q.mu.Unlock()
	q.cond.Broadcast()
}
