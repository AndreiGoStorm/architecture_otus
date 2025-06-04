package hw2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMoveObject(t *testing.T) {
	t.Run("move object to new position", func(t *testing.T) {
		// Arrange
		obj := NewSpaceship()
		obj.setProperty("position", &Vector{12, 5})
		obj.setProperty("velocity", &Vector{-7, 3})
		movable := NewMovableObjectAdapter(obj)

		move := NewMove(movable)

		// Act
		err := move.Execute()

		// Assert
		require.Nil(t, err)
		position, _ := movable.getPosition()
		require.Equal(t, position.x, float64(5))
		require.Equal(t, position.y, float64(8))
	})

	t.Run("can not read position", func(t *testing.T) {
		// Arrange
		movable := NewMovableObjectAdapter(NewSpaceship())
		move := NewMove(movable)

		// Act
		err := move.Execute()

		// Assert
		require.ErrorIs(t, err, ErrNoPosition)
	})

	t.Run("can not read velocity", func(t *testing.T) {
		// Arrange
		obj := NewSpaceship()
		obj.setProperty("position", &Vector{10, -2})
		movable := NewMovableObjectAdapter(obj)
		move := NewMove(movable)

		// Act
		err := move.Execute()

		// Assert
		require.ErrorIs(t, err, ErrNoVelocity)
	})
}
