package hw2

type Rotate struct {
	obj Rotatable
}

func NewRotate(obj Rotatable) *Rotate {
	return &Rotate{obj}
}

func (r *Rotate) Execute() error {
	direction, err := r.obj.getDirection()
	if err != nil {
		return err
	}

	angularVelocity, err := r.obj.getAngularVelocity()
	if err != nil {
		return err
	}

	directionsNumber, err := r.obj.getDirectionsNumber()
	if err != nil {
		return err
	}

	r.obj.setDirection((direction + angularVelocity) % directionsNumber)
	return nil
}
