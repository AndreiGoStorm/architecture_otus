package hw3

import (
	"testing"
)

func TestRunner(t *testing.T) {
	t.Run("run all commands", func(_ *testing.T) {
		// Arrange
		r := NewRunner()
		r.q.Enqueue(&RotateCommand{})
		r.q.Enqueue(&MoveCommand{})
		r.q.Enqueue(&FireCommand{})

		// Act
		r.run()
	})
}
