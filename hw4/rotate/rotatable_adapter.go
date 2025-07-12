package rotate

import (
	"architecture_otus/hw4"
)

type Rotatable interface {
	getAngle() (float64, error)
	setAngle(float64)
	getRotationAngle() (float64, error)
}

type RotatableObjectAdapter struct {
	obj hw4.SpaceObject
}

func NewRotatableObjectAdapter(obj hw4.SpaceObject) Rotatable {
	return &RotatableObjectAdapter{obj}
}

func (r *RotatableObjectAdapter) getAngle() (float64, error) {
	angle := r.obj.GetProperty("angle")
	if angle == nil {
		return 0, hw4.ErrNoAngle
	}
	return angle.(float64), nil
}

func (r *RotatableObjectAdapter) setAngle(angle float64) {
	r.obj.SetProperty("angle", angle)
}

func (r *RotatableObjectAdapter) getRotationAngle() (float64, error) {
	angle := r.obj.GetProperty("rotationAngle")
	if angle == nil {
		return 0, hw4.ErrNoRotationAngle
	}
	return angle.(float64), nil
}
