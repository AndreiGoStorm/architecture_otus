package hw7

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Server struct {
	q         *Queue
	wg        sync.WaitGroup
	behaviour func() bool
	stop      int32
}

func NewServer(q *Queue) *Server {
	server := &Server{q: q}
	server.behaviour = server.defaultBehaviour()

	return server
}

func (s *Server) updateBehaviour(behaviour func() bool) {
	s.behaviour = behaviour
}

func (s *Server) softBehaviour() func() bool {
	return func() bool {
		if s.q.Size() == 0 {
			s.Stop()
			return false
		}
		return s.defaultBehaviour()()
	}
}

func (s *Server) defaultBehaviour() func() bool {
	return func() bool {
		cmd := s.q.Dequeue()
		if cmd == nil {
			return false
		}
		if err := cmd.Execute(); err != nil {
			fmt.Printf("[command] error: %v\n", err)
		}
		return true
	}
}

func (s *Server) Start() {
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		for atomic.LoadInt32(&s.stop) == 0 {
			if !s.behaviour() {
				break
			}
		}
	}()
}

func (s *Server) Stop() {
	atomic.StoreInt32(&s.stop, 1)
	s.q.Stop()
	s.wg.Wait()
}
