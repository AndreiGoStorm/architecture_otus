package burnfuel

import (
	"testing"

	"architecture_otus/hw4"
	"github.com/stretchr/testify/require"
)

func TestBurnFuel(t *testing.T) {
	t.Run("burn needed fuel", func(t *testing.T) {
		// Arrange
		obj := hw4.NewSpaceship()
		obj.SetProperty("neededFuel", float64(20))
		obj.SetProperty("fuel", float64(25))
		adapter := NewBurnableFuelAdapter(obj)

		burn := NewBurnFuel(adapter)

		// Act
		err := burn.Execute()

		// Assert
		require.Nil(t, err)
		fuel := obj.GetProperty("neededFuel")
		require.Nil(t, fuel)
		fuel = obj.GetProperty("fuel")
		require.Equal(t, fuel.(float64), float64(20))
	})

	t.Run("err no needed fuel", func(t *testing.T) {
		// Arrange
		obj := hw4.NewSpaceship()
		obj.SetProperty("fuel", float64(25))
		adapter := NewBurnableFuelAdapter(obj)

		burn := NewBurnFuel(adapter)

		// Act
		err := burn.Execute()

		// Assert
		require.ErrorIs(t, err, hw4.ErrNoNeededFuel)
	})
}
