package move

import (
	"architecture_otus/hw4"
	"architecture_otus/pkg"
)

type Movable interface {
	getPosition() (*pkg.Vector, error)
	setPosition(*pkg.Vector)
	getVelocity() (*pkg.Vector, error)
}

type MovableObjectAdapter struct {
	obj hw4.SpaceObject
}

func NewMovableObjectAdapter(obj hw4.SpaceObject) Movable {
	return &MovableObjectAdapter{obj}
}

func (a *MovableObjectAdapter) getPosition() (*pkg.Vector, error) {
	position := a.obj.GetProperty("position")
	if position == nil {
		return nil, hw4.ErrNoPosition
	}
	return position.(*pkg.Vector), nil
}

func (a *MovableObjectAdapter) setPosition(position *pkg.Vector) {
	a.obj.SetProperty("position", position)
}

func (a *MovableObjectAdapter) getVelocity() (*pkg.Vector, error) {
	velocity := a.obj.GetProperty("velocity")
	if velocity == nil {
		return nil, hw4.ErrNoVelocity
	}

	return velocity.(*pkg.Vector), nil
}
