package hw4

import "errors"

var ErrDistanceLessOrEqualZero = errors.New("distance less or equal zero")
var ErrNotEnoughFuel = errors.New("not enough fuel")
var ErrNoAngle = errors.New("cannot read angle from properties")
var ErrNoNeededFuel = errors.New("cannot read needed fuel from properties")
var ErrNoPosition = errors.New("cannot read position from properties")
var ErrNoRotationAngle = errors.New("cannot read rotation angle from properties")
var ErrNoVelocity = errors.New("cannot read velocity from properties")
