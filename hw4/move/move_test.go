package move

import (
	"testing"

	"architecture_otus/hw4"
	"architecture_otus/pkg"
	"github.com/stretchr/testify/require"
)

func TestMoveObject(t *testing.T) {
	t.Run("move object to new position", func(t *testing.T) {
		// Arrange
		obj := hw4.NewSpaceship()
		obj.SetProperty("position", &pkg.Vector{X: 12, Y: 5})
		obj.SetProperty("velocity", &pkg.Vector{X: -7, Y: 3})
		adapter := NewMovableObjectAdapter(obj)

		move := NewMove(adapter)

		// Act
		err := move.Execute()

		// Assert
		require.Nil(t, err)
		position, _ := adapter.getPosition()
		require.Equal(t, position.X, float64(5))
		require.Equal(t, position.Y, float64(8))
	})

	t.Run("can not read position", func(t *testing.T) {
		// Arrange
		adapter := NewMovableObjectAdapter(hw4.NewSpaceship())
		move := NewMove(adapter)

		// Act
		err := move.Execute()

		// Assert
		require.ErrorIs(t, err, hw4.ErrNoPosition)
	})

	t.Run("can not read velocity", func(t *testing.T) {
		// Arrange
		obj := hw4.NewSpaceship()
		obj.SetProperty("position", &pkg.Vector{X: 10, Y: -2})
		adapter := NewMovableObjectAdapter(obj)
		move := NewMove(adapter)

		// Act
		err := move.Execute()

		// Assert
		require.ErrorIs(t, err, hw4.ErrNoVelocity)
	})
}
