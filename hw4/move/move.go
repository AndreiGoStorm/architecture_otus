package move

type Move struct {
	obj Movable
}

func NewMove(obj Movable) *Move {
	return &Move{obj}
}

func (m *Move) Execute() error {
	position, err := m.obj.getPosition()
	if err != nil {
		return err
	}
	velocity, err := m.obj.getVelocity()
	if err != nil {
		return err
	}

	m.obj.setPosition(position.Sum(velocity))
	return nil
}
