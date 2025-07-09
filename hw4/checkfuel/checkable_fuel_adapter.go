package checkfuel

import (
	"errors"

	"architecture_otus/hw4"
	"architecture_otus/pkg"
)

var ErrNoFuel = errors.New("cannot read fuel from properties")

type CheckableFuel interface {
	getPosition() (*pkg.Vector, error)
	getVelocity() (*pkg.Vector, error)
	getFuel() (float64, error)
	setNeededFuel(fuel float64)
}

type CheckableFuelAdapter struct {
	obj hw4.SpaceObject
}

func NewCheckableFuelAdapter(obj hw4.SpaceObject) CheckableFuel {
	return &CheckableFuelAdapter{obj}
}

func (a *CheckableFuelAdapter) getPosition() (*pkg.Vector, error) {
	position := a.obj.GetProperty("position")
	if position == nil {
		return nil, hw4.ErrNoPosition
	}
	return position.(*pkg.Vector), nil
}

func (a *CheckableFuelAdapter) getVelocity() (*pkg.Vector, error) {
	velocity := a.obj.GetProperty("velocity")
	if velocity == nil {
		return nil, hw4.ErrNoVelocity
	}

	return velocity.(*pkg.Vector), nil
}

func (a *CheckableFuelAdapter) getFuel() (float64, error) {
	fuel := a.obj.GetProperty("fuel")
	if fuel == nil {
		return 0, ErrNoFuel
	}
	return fuel.(float64), nil
}

func (a *CheckableFuelAdapter) setNeededFuel(fuel float64) {
	a.obj.SetProperty("neededFuel", fuel)
}
