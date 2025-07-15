package hw7

import "fmt"

type HardStopCommand struct {
	s *Server
}

func NewHardStopCommand(s *Server) *HardStopCommand {
	return &HardStopCommand{s}
}

func (cmd *HardStopCommand) Execute() error {
	fmt.Println("hard stop command execute")
	cmd.s.Stop()
	return nil
}
