package hw3

import (
	"fmt"
)

type MoveCommand struct{}

func (mc *MoveCommand) execute() error {
	fmt.Println("move command execute")
	return nil
}

type RotateCommand struct{}

func (rc *RotateCommand) execute() error {
	fmt.Println("rotate command execute")
	return nil
}

type FireCommand struct{}

func (rc *FireCommand) execute() error {
	fmt.Println("fire command execute")
	return nil
}
