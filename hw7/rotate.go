package hw7

import "fmt"

type RotateCommand struct{}

func (c *RotateCommand) Execute() error {
	fmt.Println("rotate command execute")
	return nil
}
