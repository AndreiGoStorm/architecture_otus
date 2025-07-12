package rotate

import (
	"testing"

	"architecture_otus/hw4"
	"github.com/stretchr/testify/require"
)

func TestRotateObject(t *testing.T) {
	t.Run("rotate object", func(t *testing.T) {
		// Arrange
		obj := hw4.NewSpaceship()
		obj.SetProperty("angle", float64(0))
		obj.SetProperty("rotationAngle", float64(90))
		adapter := NewRotatableObjectAdapter(obj)

		rotate := NewRotate(adapter)

		// Act
		err := rotate.Execute()

		// Assert
		require.Nil(t, err)
		angle, _ := adapter.getAngle()
		require.Equal(t, angle, float64(90))
	})

	t.Run("can not read angle", func(t *testing.T) {
		// Arrange
		adapter := NewRotatableObjectAdapter(hw4.NewSpaceship())
		rotate := NewRotate(adapter)

		// Act
		err := rotate.Execute()

		// Assert
		require.ErrorIs(t, err, hw4.ErrNoAngle)
	})

	t.Run("can not read rotationAngle", func(t *testing.T) {
		// Arrange
		obj := hw4.NewSpaceship()
		obj.SetProperty("angle", float64(50))
		adapter := NewRotatableObjectAdapter(obj)
		rotate := NewRotate(adapter)

		// Act
		err := rotate.Execute()

		// Assert
		require.ErrorIs(t, err, hw4.ErrNoRotationAngle)
	})
}
