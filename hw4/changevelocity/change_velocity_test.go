package changevelocity

import (
	"testing"

	"architecture_otus/hw4"
	"architecture_otus/pkg"
	"github.com/stretchr/testify/require"
)

func TestChangeVelocity(t *testing.T) {
	t.Run("change velocity", func(t *testing.T) {
		// Arrange
		obj := hw4.NewSpaceship()
		obj.SetProperty("velocity", &pkg.Vector{X: 5, Y: 0})
		obj.SetProperty("rotationAngle", 90.0)
		adapter := NewChangeableVelocityAdapter(obj)
		change := NewChangeVelocity(adapter)

		// Act
		err := change.Execute()

		// Assert
		require.Nil(t, err)
		velocity, err := adapter.getVelocity()
		require.Nil(t, err)
		require.InDelta(t, velocity.X, 0, hw4.Delta)
		require.InDelta(t, velocity.Y, 5, hw4.Delta)
	})

	t.Run("no change velocity for stationary object", func(t *testing.T) {
		// Arrange
		obj := hw4.NewSpaceship()
		obj.SetProperty("velocity", &pkg.Vector{X: 0, Y: 0})
		obj.SetProperty("rotationAngle", 90.0)
		adapter := NewChangeableVelocityAdapter(obj)
		change := NewChangeVelocity(adapter)

		// Act
		err := change.Execute()

		// Assert
		require.Nil(t, err)
		velocity, err := adapter.getVelocity()
		require.Nil(t, err)
		require.Equal(t, velocity.X, float64(0))
		require.Equal(t, velocity.Y, float64(0))
	})

	t.Run("cannot read velocity", func(t *testing.T) {
		// Arrange
		obj := hw4.NewSpaceship()
		adapter := NewChangeableVelocityAdapter(obj)
		change := NewChangeVelocity(adapter)

		// Act
		err := change.Execute()

		// Assert
		// Assert
		require.ErrorIs(t, err, hw4.ErrNoVelocity)
	})

	t.Run("cannot read rotation angle", func(t *testing.T) {
		// Arrange
		obj := hw4.NewSpaceship()
		obj.SetProperty("velocity", &pkg.Vector{X: 5, Y: 0})
		adapter := NewChangeableVelocityAdapter(obj)
		change := NewChangeVelocity(adapter)

		// Act
		err := change.Execute()

		// Assert
		// Assert
		require.ErrorIs(t, err, hw4.ErrNoRotationAngle)
	})
}
