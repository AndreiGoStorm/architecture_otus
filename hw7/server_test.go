package hw7

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
		q.Enqueue(&MoveCommand{})
		q.Enqueue(&RotateCommand{})

		// Assert
		time.Sleep(time.Second)
		server.Stop()
		require.Equal(t, 0, q.Size())
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
