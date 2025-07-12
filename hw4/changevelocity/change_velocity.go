package changevelocity

import (
	"math"

	"architecture_otus/pkg"
)

type ChangeVelocity struct {
	obj ChangeableVelocity
}

func NewChangeVelocity(obj ChangeableVelocity) *ChangeVelocity {
	return &ChangeVelocity{obj}
}

func (cv *ChangeVelocity) Execute() error {
	velocity, err := cv.obj.getVelocity()
	if err != nil {
		return err
	}
	if velocity.Length() == 0 {
		return nil
	}
	angle, err := cv.obj.getRotationAngle()
	if err != nil {
		return err
	}
	rad := angle * math.Pi / 180.0
	cos := math.Cos(rad)
	sin := math.Sin(rad)
	cv.obj.setVelocity(&pkg.Vector{
		X: velocity.X*cos - velocity.Y*sin,
		Y: velocity.X*sin + velocity.Y*cos,
	})
	return nil
}
