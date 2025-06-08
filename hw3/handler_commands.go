package hw3

import (
	"fmt"
)

type LogCommand struct {
	cmd Command
	err error
}

func (lc *LogCommand) execute() error {
	fmt.Printf("Err: %v for command %s\n", lc.err, getTypeCmd(lc.cmd))
	return nil
}

type RepeatCommand struct {
	cmd Command
	err error
	q   *Queue
}

func (rc *RepeatCommand) execute() error {
	fmt.Printf("RepeatCommand for command %s\n", getTypeCmd(rc.cmd))
	rc.q.Enqueue(rc.cmd)
	return nil
}

type RepeatLogCommand struct {
	cmd   Command
	err   error
	q     *Queue
	tries int
}

func (rlc *RepeatLogCommand) addParams(cmd Command, err error, q *Queue) {
	rlc.cmd = cmd
	rlc.err = err
	rlc.q = q
}

func (rlc *RepeatLogCommand) execute() error {
	defer func() {
		rlc.tries--
	}()

	if rlc.tries > 0 {
		fmt.Printf("RepeatLogCommand for RepeatCommand, tries %d\n", rlc.tries)
		rc := &RepeatCommand{rlc.cmd, rlc.err, rlc.q}
		return rc.execute()
	}

	fmt.Printf("RepeatLogCommand for LogCommand, tries %d\n", rlc.tries)
	lc := &LogCommand{rlc.cmd, rlc.err}
	return lc.execute()
}
