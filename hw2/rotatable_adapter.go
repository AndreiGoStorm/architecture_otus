package hw2

import "errors"

var (
	ErrNoDirection        = errors.New("cannot read direction from properties")
	ErrNoAngularVelocity  = errors.New("cannot read angular velocity from properties")
	ErrNoDirectionsNumber = errors.New("cannot read directions number from properties")
)

type Rotatable interface {
	getDirection() (int, error)
	setDirection(int)
	getAngularVelocity() (int, error)
	getDirectionsNumber() (int, error)
}

type RotatableObjectAdapter struct {
	obj SpaceObject
}

func NewRotatableObjectAdapter(obj SpaceObject) Rotatable {
	return &RotatableObjectAdapter{obj}
}

func (r *RotatableObjectAdapter) getDirection() (int, error) {
	direction := r.obj.getProperty("direction")
	if direction == nil {
		return 0, ErrNoDirection
	}
	return direction.(int), nil
}

func (r *RotatableObjectAdapter) setDirection(direction int) {
	r.obj.setProperty("direction", direction)
}

func (r *RotatableObjectAdapter) getAngularVelocity() (int, error) {
	angularVelocity := r.obj.getProperty("angularVelocity")
	if angularVelocity == nil {
		return 0, ErrNoAngularVelocity
	}
	return angularVelocity.(int), nil
}

func (r *RotatableObjectAdapter) getDirectionsNumber() (int, error) {
	directionsNumber := r.obj.getProperty("directionsNumber")
	if directionsNumber == nil {
		return 0, ErrNoDirectionsNumber
	}
	return directionsNumber.(int), nil
}
