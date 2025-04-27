package hw1

import (
	"crypto/rand"
	"encoding/binary"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

var a, b, c float64

func TestQuadraticEquation(t *testing.T) {
	t.Run("equation has no roots", func(t *testing.T) {
		// Arrange
		a = 1
		b = 0
		c = 1

		// Act
		_, _, err := solve(a, b, c)

		// Assert
		require.Error(t, err)
		require.ErrorIs(t, err, ErrNoRoots)
	})

	t.Run("equation has one root", func(t *testing.T) {
		// Arrange
		a = 1
		b = 2
		c = 1
		expectedX := -1

		// Act
		x1, x2, err := solve(a, b, c)

		require.NoError(t, err)
		require.InDelta(t, expectedX, x1, delta)
		require.InDelta(t, expectedX, x2, delta)
	})

	t.Run("equation has two roots", func(t *testing.T) {
		// Arrange
		a = 1
		b = 0
		c = -1
		expectedX1 := 1
		expectedX2 := -1

		// Act
		x1, x2, err := solve(a, b, c)

		// Assert
		require.NoError(t, err)
		require.InDelta(t, expectedX1, x1, delta)
		require.InDelta(t, expectedX2, x2, delta)
	})

	t.Run("error a equal zero", func(t *testing.T) {
		// Arrange
		a = 0
		b, _ = cryptoRandomFloat64()
		c, _ = cryptoRandomFloat64()

		// Act
		_, _, err := solve(a, b, c)

		// Assert
		require.Error(t, err)
		require.ErrorIs(t, err, ErrAEqualZero)
	})
}

func cryptoRandomFloat64() (float64, error) {
	var buf [8]byte
	_, err := rand.Read(buf[:])
	if err != nil {
		return 0, err
	}

	randomUint64 := binary.LittleEndian.Uint64(buf[:])

	base := float64(randomUint64) / float64(math.MaxUint64)
	return base * 20, nil
}
