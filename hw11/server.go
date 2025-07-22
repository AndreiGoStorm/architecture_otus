package hw11

import (
	"fmt"
	"sync"
)

type Server struct {
	q         *Queue
	moteToQ   *Queue
	stop      chan struct{}
	behaviour func() bool
	wg        sync.WaitGroup
	once      sync.Once

	st *ServerState
}

func NewServer(q *Queue) *Server {
	server := &Server{q: q, stop: make(chan struct{})}
	server.moteToQ = NewQueue()
	server.behaviour = server.defaultBehaviour()
	server.st = NewServerState(server)

	return server
}

func (s *Server) Start() {
	s.wg.Add(1)
	go s.worker()
}

func (s *Server) worker() {
	defer s.wg.Done()
	for {
		select {
		/** заблокироваться на чтение из канала stop и ждать, пока в нём появится значение или пока канал не закроется */
		case <-s.stop:
			fmt.Println("worker: stopped")
			return
		default:
			fmt.Println("worker: working...")
			if !s.behaviour() {
				return
			}
		}
	}
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
		return s.st.run()
	}
}

func (s *Server) Stop() {
	s.once.Do(func() {
		close(s.stop)
		s.q.Stop()
		s.moteToQ.Stop()
	})
}
