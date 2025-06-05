package hw2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRotateObject(t *testing.T) {
	t.Run("rotate object to new position", func(t *testing.T) {
		// Arrange
		obj := NewSpaceship()
		obj.setProperty("direction", 25)
		obj.setProperty("angularVelocity", 120)
		obj.setProperty("directionsNumber", 8)
		rotatable := NewRotatableObjectAdapter(obj)

		rotate := NewRotate(rotatable)

		// Act
		err := rotate.Execute()

		// Assert
		require.Nil(t, err)
		direction, _ := rotatable.getDirection()
		require.Equal(t, direction, 1)
	})

	t.Run("can not read direction", func(t *testing.T) {
		// Arrange
		rotatable := NewRotatableObjectAdapter(NewSpaceship())
		rotate := NewRotate(rotatable)

		// Act
		err := rotate.Execute()

		// Assert
		require.ErrorIs(t, err, ErrNoDirection)
	})

	t.Run("can not read angular velocity", func(t *testing.T) {
		// Arrange
		obj := NewSpaceship()
		obj.setProperty("direction", 10)
		rotatable := NewRotatableObjectAdapter(obj)
		rotate := NewRotate(rotatable)

		// Act
		err := rotate.Execute()

		// Assert
		require.ErrorIs(t, err, ErrNoAngularVelocity)
	})

	t.Run("can not read directions number", func(t *testing.T) {
		// Arrange
		obj := NewSpaceship()
		obj.setProperty("direction", 10)
		obj.setProperty("angularVelocity", 2)
		rotatable := NewRotatableObjectAdapter(obj)
		rotate := NewRotate(rotatable)

		// Act
		err := rotate.Execute()

		// Assert
		require.ErrorIs(t, err, ErrNoDirectionsNumber)
	})
}
