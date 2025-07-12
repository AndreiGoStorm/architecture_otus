package burnfuel

import (
	"architecture_otus/hw4"
)

type BurnableFuel interface {
	getNeededFuel() (float64, error)
	setFuel(fuel float64)
}

type BurnableFuelAdapter struct {
	obj hw4.SpaceObject
}

func NewBurnableFuelAdapter(obj hw4.SpaceObject) BurnableFuel {
	return &BurnableFuelAdapter{obj}
}

func (a *BurnableFuelAdapter) getNeededFuel() (float64, error) {
	fuel := a.obj.GetProperty("neededFuel")
	if fuel == nil {
		return 0, hw4.ErrNoNeededFuel
	}
	return fuel.(float64), nil
}

func (a *BurnableFuelAdapter) setFuel(fuel float64) {
	a.obj.SetProperty("fuel", fuel)
	a.obj.SetProperty("neededFuel", nil)
}
