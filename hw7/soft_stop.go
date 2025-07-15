package hw7

import "fmt"

type SoftStopCommand struct {
	s *Server
}

func NewSoftStopCommand(s *Server) *SoftStopCommand {
	return &SoftStopCommand{s}
}

func (cmd *SoftStopCommand) Execute() error {
	fmt.Println("soft stop command execute")
	cmd.s.updateBehaviour(cmd.s.softBehaviour())
	return nil
}
