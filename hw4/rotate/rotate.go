package rotate

import "math"

type Rotate struct {
	obj Rotatable
}

func NewRotate(obj Rotatable) *Rotate {
	return &Rotate{obj}
}

func (r *Rotate) Execute() error {
	angle, err := r.obj.getAngle()
	if err != nil {
		return err
	}
	rotationAngle, err := r.obj.getRotationAngle()
	if err != nil {
		return err
	}
	r.obj.setAngle(math.Mod(angle+rotationAngle, 360))

	return nil
}
