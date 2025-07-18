package hw2

import (
	"errors"

	"architecture_otus/pkg"
)

var (
	ErrNoPosition = errors.New("cannot read position from properties")
	ErrNoVelocity = errors.New("cannot read velocity from properties")
)

type Movable interface {
	getPosition() (*pkg.Vector, error)
	setPosition(*pkg.Vector)
	getVelocity() (*pkg.Vector, error)
}

type MovableObjectAdapter struct {
	obj SpaceObject
}

func NewMovableObjectAdapter(obj SpaceObject) Movable {
	return &MovableObjectAdapter{obj}
}

func (a *MovableObjectAdapter) getPosition() (*pkg.Vector, error) {
	position := a.obj.getProperty("position")
	if position == nil {
		return nil, ErrNoPosition
	}
	return position.(*pkg.Vector), nil
}

func (a *MovableObjectAdapter) setPosition(position *pkg.Vector) {
	a.obj.setProperty("position", position)
}

func (a *MovableObjectAdapter) getVelocity() (*pkg.Vector, error) {
	velocity := a.obj.getProperty("velocity")
	if velocity == nil {
		return nil, ErrNoVelocity
	}

	return velocity.(*pkg.Vector), nil
}
