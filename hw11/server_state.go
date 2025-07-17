package hw11

type ServerState struct {
	current State

	simple State
	moveTo State
}

func NewServerState(s *Server) *ServerState {
	st := &ServerState{}
	st.initStates(s)
	st.setSimpleState()
	return st
}

func (st *ServerState) initStates(s *Server) {
	st.simple = &SimpleState{s}
	st.moveTo = &MoveToState{s}
}

func (st *ServerState) setSimpleState() {
	st.current = st.simple
}

func (st *ServerState) setMoveToState() {
	st.current = st.moveTo
}

func (st *ServerState) run() bool {
	return st.current.handle()
}
