package hw11

import "fmt"

type Command interface {
	Execute() error
}

type MoveCommand struct{}

func (c *MoveCommand) Execute() error {
	fmt.Println("move command execute")
	return nil
}

type RotateCommand struct{}

func (c *RotateCommand) Execute() error {
	fmt.Println("rotate command execute")
	return nil
}

type ServerCommand interface {
	getServer() *Server
}

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

func (cmd *HardStopCommand) getServer() *Server {
	return cmd.s
}

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

func (cmd *SoftStopCommand) getServer() *Server {
	return cmd.s
}

type RunCommand struct {
	s *Server
}

func NewRunCommand(s *Server) *RunCommand {
	return &RunCommand{s}
}

func (cmd *RunCommand) Execute() error {
	fmt.Println("run command execute")
	cmd.s.st.setSimpleState()
	return nil
}

func (cmd *RunCommand) getServer() *Server {
	return cmd.s
}

type MoveToCommand struct {
	s *Server
}

func NewMoveToCommand(s *Server) *MoveToCommand {
	return &MoveToCommand{s}
}

func (cmd *MoveToCommand) Execute() error {
	fmt.Println("move to command execute")
	cmd.s.st.setMoveToState()
	return nil
}

func (cmd *MoveToCommand) getServer() *Server {
	return cmd.s
}
