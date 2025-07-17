package hw11

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestServerThread(t *testing.T) {
	t.Run("server start", func(t *testing.T) {
		// Arrange
		q := NewQueue()
		server := NewServer(q)

		// Act
		server.Start()
		server.q.Enqueue(new(MoveCommand))
		server.q.Enqueue(&RotateCommand{})

		// Assert
		time.Sleep(time.Second)
		server.Stop()
		require.Equal(t, 0, server.q.Size())
	})

	t.Run("server stop", func(t *testing.T) {
		// Arrange
		q := NewQueue()
		server := NewServer(q)

		// Act
		server.Start()
		q.Enqueue(&MoveCommand{})
		done := make(chan struct{})
		go func() {
			server.Stop()
			close(done)
		}()

		// Assert
		select {
		case <-done:
		case <-time.After(time.Second):
			t.Fatal("Server not respond")
		}
	})
}

func TestServerStop(t *testing.T) {
	t.Run("server simple stop", func(t *testing.T) {
		// Arrange
		q := NewQueue()
		server := NewServer(q)
		server.Start()

		// Act
		q.Enqueue(&MoveCommand{})
		q.Enqueue(&RotateCommand{})

		// Assert
		server.Stop()
		require.Equal(t, 2, q.Size())
	})

	t.Run("server soft stop", func(t *testing.T) {
		// Arrange
		q := NewQueue()
		server := NewServer(q)
		server.Start()

		// Act
		q.Enqueue(&MoveCommand{})
		q.Enqueue(NewSoftStopCommand(server))
		q.Enqueue(&RotateCommand{})

		// Assert
		time.Sleep(time.Second)
		require.Equal(t, 0, q.Size())
	})

	t.Run("server hard stop", func(t *testing.T) {
		// Arrange
		q := NewQueue()
		server := NewServer(q)
		server.Start()

		// Act
		q.Enqueue(&MoveCommand{})
		q.Enqueue(&RotateCommand{})
		q.Enqueue(NewHardStopCommand(server))
		q.Enqueue(&MoveCommand{})
		q.Enqueue(&RotateCommand{})

		// Assert
		time.Sleep(time.Second)
		require.Equal(t, 2, q.Size())
	})
}

func TestSimpleState(t *testing.T) {
	t.Run("server simple state", func(t *testing.T) {
		// Arrange
		q := NewQueue()
		server := NewServer(q)

		// Act
		server.Start()
		server.q.Enqueue(&MoveCommand{})
		server.q.Enqueue(&RotateCommand{})
		server.q.Enqueue(NewMoveToCommand(server))
		server.q.Enqueue(&RotateCommand{})
		server.q.Enqueue(&MoveCommand{})

		// Assert
		time.Sleep(time.Second)
		server.Stop()
		require.Equal(t, 0, server.q.Size())
		require.Equal(t, 2, server.moteToQ.Size())
		require.IsType(t, &MoveToState{}, server.st.current)
	})

	t.Run("server simple state hard stop", func(t *testing.T) {
		// Arrange
		q := NewQueue()
		server := NewServer(q)

		// Act
		server.Start()
		server.q.Enqueue(&MoveCommand{})
		server.q.Enqueue(NewMoveToCommand(server))
		server.q.Enqueue(&RotateCommand{})
		server.q.Enqueue(&MoveCommand{})
		server.q.Enqueue(&RotateCommand{})
		server.q.Enqueue(NewHardStopCommand(server))
		server.q.Enqueue(&MoveCommand{})

		// Assert
		time.Sleep(time.Second)
		server.Stop()
		require.Equal(t, 1, server.q.Size())
		require.Equal(t, 3, server.moteToQ.Size())
		require.IsType(t, &MoveToState{}, server.st.current)
	})
}

func TestMoveToState(t *testing.T) {
	t.Run("server move to state", func(t *testing.T) {
		// Arrange
		q := NewQueue()
		server := NewServer(q)
		server.st.setMoveToState()

		// Act
		server.Start()
		server.q.Enqueue(&MoveCommand{})
		server.q.Enqueue(&RotateCommand{})
		server.q.Enqueue(NewRunCommand(server))
		server.q.Enqueue(&RotateCommand{})
		server.q.Enqueue(&MoveCommand{})

		// Assert
		time.Sleep(time.Second * 1)
		server.Stop()
		require.Equal(t, 0, server.q.Size())
		require.Equal(t, 2, server.moteToQ.Size())
		require.IsType(t, &SimpleState{}, server.st.current)
	})

	t.Run("server move to change state", func(t *testing.T) {
		// Arrange
		q := NewQueue()
		server := NewServer(q)
		server.st.setMoveToState()

		// Act
		server.Start()
		server.q.Enqueue(&RotateCommand{})
		server.q.Enqueue(NewRunCommand(server))
		server.q.Enqueue(&RotateCommand{})
		server.q.Enqueue(&MoveCommand{})
		server.q.Enqueue(NewMoveToCommand(server))
		server.q.Enqueue(&MoveCommand{})

		// Assert
		time.Sleep(time.Second)
		server.Stop()
		require.Equal(t, 0, server.q.Size())
		require.Equal(t, 2, server.moteToQ.Size())
		require.IsType(t, &MoveToState{}, server.st.current)
	})

	t.Run("server move to hard stop", func(t *testing.T) {
		// Arrange
		q := NewQueue()
		server := NewServer(q)
		server.st.setMoveToState()

		// Act
		server.Start()
		server.q.Enqueue(&RotateCommand{})
		server.q.Enqueue(&MoveCommand{})
		server.q.Enqueue(NewHardStopCommand(server))
		server.q.Enqueue(&MoveCommand{})
		server.q.Enqueue(&MoveCommand{})
		server.q.Enqueue(&RotateCommand{})

		// Assert
		time.Sleep(time.Second)
		server.Stop()
		require.Equal(t, 3, server.q.Size())
		require.Equal(t, 2, server.moteToQ.Size())
		require.IsType(t, &MoveToState{}, server.st.current)
	})
}
