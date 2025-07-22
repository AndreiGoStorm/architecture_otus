package hw11

import "fmt"

type State interface {
	handle() bool
}

type SimpleState struct {
	s *Server
}

func (st *SimpleState) handle() bool {
	cmd := st.s.q.Dequeue()
	if cmd == nil {
		fmt.Println("queue closed, exiting worker")
		return false
	}
	if err := cmd.Execute(); err != nil {
		fmt.Printf("command error: %v\n", err)
	}
	return true
}

type MoveToState struct {
	s *Server
}

func (st *MoveToState) handle() bool {
	cmd := st.s.q.Dequeue()
	if cmd == nil {
		fmt.Println("queue closed, exiting worker")
		return false
	}
	if _, ok := cmd.(ServerCommand); ok {
		if err := cmd.Execute(); err != nil {
			fmt.Printf("command error: %v\n", err)
		}
	} else {
		st.s.moteToQ.Enqueue(cmd)
	}
	return true
}
