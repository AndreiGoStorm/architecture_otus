package hw3

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type MockMoveCommand struct {
	mock.Mock
	Err error
}

func (m *MockMoveCommand) execute() error {
	return m.Err
}

type MockRotateCommand struct {
	mock.Mock
	Err error
}

func (m *MockRotateCommand) execute() error {
	return m.Err
}

type MockRepeatCommand struct {
	mock.Mock
	Err error
}

func (m *MockRepeatCommand) execute() error {
	return m.Err
}

type MockLogCommand struct {
	mock.Mock
	Err error
}

func (m *MockLogCommand) execute() error {
	return m.Err
}

type MockRepeatLogCommand struct {
	mock.Mock
	Err error
}

func (m *MockRepeatLogCommand) execute() error {
	return m.Err
}

func TestRepeatCommand(t *testing.T) {
	t.Run("repeat command when move command throw exception", func(t *testing.T) {
		// Arrange
		r := NewRunner()
		cmd := &MoveCommand{}

		// Act
		command := r.eh.handle(cmd, errors.New("ErrMoveRepeatHandler"))
		err := command.execute()

		// Assert
		require.NoError(t, err)
		require.Equal(t, r.q.Len(), 1)
		require.Equal(t, getTypeCmd(command), "RepeatCommand")
	})

	t.Run("repeat when repeat command throw exception", func(t *testing.T) {
		// Arrange
		r := NewRunner()
		cmd := &MockMoveCommand{Err: errors.New("ErrMoveRepeatHandler")}
		repeatHandler := func(_ Command, _ error) Command {
			return &MockRepeatCommand{Err: errors.New("ErrRepeatCommand")}
		}
		r.eh.register(getTypeCmd(cmd), cmd.Err.Error(), repeatHandler)

		// Act
		r.execute(cmd)

		// Assert
		require.Equal(t, r.q.Len(), 1)
		elem := r.q.Dequeue()
		command := elem.(Command)
		require.Equal(t, getTypeCmd(command), "MockRepeatCommand")
	})
}

func TestLogCommand(t *testing.T) {
	t.Run("write exception to log when rotate command throw exception", func(t *testing.T) {
		// Arrange
		r := NewRunner()
		cmd := &RotateCommand{}

		// Act
		command := r.eh.handle(cmd, errors.New("ErrRotateLogHandler"))
		err := command.execute()

		// Assert
		require.NoError(t, err)
		require.Equal(t, r.q.Len(), 0)
		require.Equal(t, getTypeCmd(command), "LogCommand")
	})
}

func TestRepeatLogCommand(t *testing.T) {
	t.Run("add repeat log command then when move command throw exception", func(t *testing.T) {
		// Arrange
		r := NewRunner()
		cmd := &MockMoveCommand{Err: errors.New("ErrRepeatLogHandler")}
		rlc := &RepeatLogCommand{tries: 1}
		repeatLogHandler := func(cmd Command, err error) Command {
			rlc.addParams(cmd, err, r.q)
			return rlc
		}
		r.eh.register("MockMoveCommand", cmd.Err.Error(), repeatLogHandler)

		// Act
		r.execute(cmd)

		// Assert
		require.Equal(t, r.q.Len(), 1)
		elem := r.q.Dequeue()
		command := elem.(Command)
		r.execute(command)
		require.Equal(t, r.q.Len(), 0)
	})

	t.Run("add repeat log command then when move command throw exception", func(t *testing.T) {
		// Arrange
		r := NewRunner()
		cmd := &MockRotateCommand{Err: errors.New("ErrRotateRepeatLogHandler")}
		rlc := &RepeatLogCommand{tries: repeatTries}
		repeatLogHandler := func(cmd Command, err error) Command {
			rlc.addParams(cmd, err, r.q)
			return rlc
		}
		r.eh.register("MockRotateCommand", cmd.Err.Error(), repeatLogHandler)

		// Act
		r.execute(cmd)

		// Assert
		require.Equal(t, r.q.Len(), 1)
		elem := r.q.Dequeue()
		command := elem.(Command)
		r.execute(command)
	})
}
