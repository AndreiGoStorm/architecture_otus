package hw7

import (
	"fmt"
)

type MoveCommand struct{}

func (c *MoveCommand) Execute() error {
	fmt.Println("move command execute")
	return nil
}
