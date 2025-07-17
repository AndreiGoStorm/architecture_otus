package hw11

import (
	"sync"
	"sync/atomic"
)

type Queue struct {
	ch    chan Command
	count int32
	once  sync.Once
	mu    sync.Mutex
}

func NewQueue() *Queue {
	return &Queue{
		ch: make(chan Command, 10),
	}
}

func (q *Queue) Enqueue(cmd Command) {
	q.mu.Lock()
	defer q.mu.Unlock()
	atomic.AddInt32(&q.count, 1)
	q.ch <- cmd
}

func (q *Queue) Dequeue() Command {
	cmd, ok := <-q.ch
	if !ok {
		return nil // канал закрыт
	}
	atomic.AddInt32(&q.count, -1)
	return cmd
}

func (q *Queue) Size() int {
	return int(atomic.LoadInt32(&q.count))
}

func (q *Queue) Stop() {
	q.once.Do(func() {
		q.mu.Lock()
		defer q.mu.Unlock()
		close(q.ch)
	})
}
