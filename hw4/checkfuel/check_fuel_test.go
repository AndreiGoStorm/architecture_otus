package checkfuel

import (
	"testing"

	"architecture_otus/hw4"
	"architecture_otus/pkg"
	"github.com/stretchr/testify/require"
)

func TestCheckFuel(t *testing.T) {
	t.Run("check needed fuel", func(t *testing.T) {
		// Arrange
		obj := hw4.NewSpaceship()
		obj.SetProperty("position", &pkg.Vector{X: 4, Y: 3})
		obj.SetProperty("velocity", &pkg.Vector{X: 4, Y: 3})
		obj.SetProperty("fuel", float64(25))
		adapter := NewCheckableFuelAdapter(obj)

		check := NewCheckFuel(adapter)

		// Act
		err := check.Execute()

		// Assert
		require.Nil(t, err)
		fuel := obj.GetProperty("neededFuel")
		require.InDelta(t, fuel.(float64), 20, hw4.Delta)
	})

	t.Run("not enough fuel", func(t *testing.T) {
		// Arrange
		obj := hw4.NewSpaceship()
		obj.SetProperty("position", &pkg.Vector{X: 6, Y: 10})
		obj.SetProperty("velocity", &pkg.Vector{X: -6, Y: 20})
		obj.SetProperty("fuel", float64(20))
		adapter := NewCheckableFuelAdapter(obj)

		check := NewCheckFuel(adapter)

		// Act
		err := check.Execute()

		// Assert
		require.ErrorIs(t, err, hw4.ErrNotEnoughFuel)
	})
}
