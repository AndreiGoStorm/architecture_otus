package hw5

type Rotatable interface {
	getAngle() (float64, error)
	setAngle(float64)
	getRotationAngle() (float64, error)
}

type RotatableObjectAdapter struct {
	obj SpaceObject
}

func NewRotatableObjectAdapter(obj SpaceObject) Rotatable {
	return &RotatableObjectAdapter{obj}
}

func (r *RotatableObjectAdapter) getAngle() (float64, error) {
	angle := r.obj.GetProperty("angle")
	if angle == nil {
		return 0, ErrNoAngle
	}
	return angle.(float64), nil
}

func (r *RotatableObjectAdapter) setAngle(angle float64) {
	r.obj.SetProperty("angle", angle)
}

func (r *RotatableObjectAdapter) getRotationAngle() (float64, error) {
	angle := r.obj.GetProperty("rotationAngle")
	if angle == nil {
		return 0, ErrNoRotationAngle
	}
	return angle.(float64), nil
}
