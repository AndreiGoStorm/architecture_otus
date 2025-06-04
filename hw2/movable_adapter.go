package hw2

import "errors"

var (
	ErrNoPosition = errors.New("cannot read position from properties")
	ErrNoVelocity = errors.New("cannot read velocity from properties")
)

type Movable interface {
	getPosition() (*Vector, error)
	setPosition(*Vector)
	getVelocity() (*Vector, error)
}

type MovableObjectAdapter struct {
	obj SpaceObject
}

func NewMovableObjectAdapter(obj SpaceObject) Movable {
	return &MovableObjectAdapter{obj}
}

func (a *MovableObjectAdapter) getPosition() (*Vector, error) {
	position := a.obj.getProperty("position")
	if position == nil {
		return nil, ErrNoPosition
	}
	return position.(*Vector), nil
}

func (a *MovableObjectAdapter) setPosition(position *Vector) {
	a.obj.setProperty("position", position)
}

func (a *MovableObjectAdapter) getVelocity() (*Vector, error) {
	velocity := a.obj.getProperty("velocity")
	if velocity == nil {
		return nil, ErrNoVelocity
	}

	return velocity.(*Vector), nil
}
