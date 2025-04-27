package hw1

import (
	"errors"
	"math"
)

var (
	ErrNoRoots    = errors.New("no roots")
	ErrAEqualZero = errors.New("coefficient a equals zero")
)

const delta = 0.000001

func solve(a, b, c float64) (x1, x2 float64, err error) {
	if math.Abs(a) <= delta {
		return 0, 0, ErrAEqualZero
	}

	D := b*b - 4*a*c
	if D < -delta {
		return 0, 0, ErrNoRoots
	}

	if math.Abs(D) <= delta {
		x1 = -b / 2 * a
		return x1, x1, nil
	}

	sqrt := math.Sqrt(D)
	return (-b + sqrt) / (2 * a), (-b - sqrt) / (2 * a), nil
}
