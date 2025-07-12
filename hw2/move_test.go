package hw2

import (
	"testing"

	"architecture_otus/pkg"
	"github.com/stretchr/testify/require"
)

func TestMoveObject(t *testing.T) {
	t.Run("move object to new position", func(t *testing.T) {
		// Arrange
		obj := NewSpaceship()
		obj.setProperty("position", &pkg.Vector{X: 12, Y: 5})
		obj.setProperty("velocity", &pkg.Vector{X: -7, Y: 3})
		movable := NewMovableObjectAdapter(obj)

		move := NewMove(movable)

		// Act
		err := move.Execute()

		// Assert
		require.Nil(t, err)
		position, _ := movable.getPosition()
		require.Equal(t, position.X, float64(5))
		require.Equal(t, position.Y, float64(8))
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
		obj.setProperty("position", &pkg.Vector{X: 10, Y: -2})
		movable := NewMovableObjectAdapter(obj)
		move := NewMove(movable)

		// Act
		err := move.Execute()

		// Assert
		require.ErrorIs(t, err, ErrNoVelocity)
	})
}
