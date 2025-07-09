package app

import (
	"testing"

	"architecture_otus/hw4"
	"architecture_otus/pkg"
	"github.com/stretchr/testify/require"
)

func TestMoveRunner(t *testing.T) {
	t.Run("run all move commands", func(_ *testing.T) {
		// Arrange
		obj := hw4.NewSpaceship()
		obj.SetProperty("position", &pkg.Vector{X: 3, Y: 4})
		obj.SetProperty("velocity", &pkg.Vector{X: -3, Y: 4})
		obj.SetProperty("fuel", float64(25))

		// Act
		err := Move(obj)

		// Assert
		require.Nil(t, err)
		fuel := obj.GetProperty("neededFuel")
		require.Nil(t, fuel)
		fuel = obj.GetProperty("fuel")
		require.Equal(t, fuel.(float64), float64(20))
	})

	t.Run("run all rotate commands", func(_ *testing.T) {
		// Arrange
		obj := hw4.NewSpaceship()
		obj.SetProperty("velocity", &pkg.Vector{X: 0, Y: 150})
		obj.SetProperty("angle", float64(-45))
		obj.SetProperty("rotationAngle", float64(90))

		// Act
		err := Rotate(obj)

		// Assert
		require.Nil(t, err)
		velocity := obj.GetProperty("velocity").(*pkg.Vector)
		require.InDelta(t, velocity.X, -150, hw4.Delta)
		require.InDelta(t, velocity.Y, 0, hw4.Delta)
		angle := obj.GetProperty("angle")
		require.InDelta(t, angle, 45, hw4.Delta)
	})
}
