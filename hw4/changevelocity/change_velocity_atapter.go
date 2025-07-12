package changevelocity

import (
	"architecture_otus/hw4"
	"architecture_otus/pkg"
)

type ChangeableVelocity interface {
	getVelocity() (*pkg.Vector, error)
	setVelocity(*pkg.Vector)
	getRotationAngle() (float64, error)
}

type ChangeableVelocityAdapter struct {
	obj hw4.SpaceObject
}

func NewChangeableVelocityAdapter(obj hw4.SpaceObject) ChangeableVelocity {
	return &ChangeableVelocityAdapter{obj}
}

func (a *ChangeableVelocityAdapter) getVelocity() (*pkg.Vector, error) {
	velocity := a.obj.GetProperty("velocity")
	if velocity == nil {
		return nil, hw4.ErrNoVelocity
	}

	return velocity.(*pkg.Vector), nil
}

func (a *ChangeableVelocityAdapter) setVelocity(velocity *pkg.Vector) {
	a.obj.SetProperty("velocity", velocity)
}

func (a *ChangeableVelocityAdapter) getRotationAngle() (float64, error) {
	angle := a.obj.GetProperty("rotationAngle")
	if angle == nil {
		return 0, hw4.ErrNoRotationAngle
	}
	return angle.(float64), nil
}
