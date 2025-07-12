package checkfuel

import (
	"math"

	"architecture_otus/hw4"
	"architecture_otus/pkg"
)

type CheckFuel struct {
	obj CheckableFuel
}

func NewCheckFuel(obj CheckableFuel) *CheckFuel {
	return &CheckFuel{obj}
}

func (cf *CheckFuel) Execute() error {
	start, err := cf.obj.getPosition()
	if err != nil {
		return err
	}
	velocity, err := cf.obj.getVelocity()
	if err != nil {
		return err
	}

	fuel, err := cf.obj.getFuel()
	if err != nil {
		return err
	}

	distance := cf.getDistance(start, start.Sum(velocity))
	if math.Abs(distance) <= hw4.Delta {
		return hw4.ErrDistanceLessOrEqualZero
	}

	neededFuel := fuel - cf.calculateFuel(distance)
	if neededFuel < -hw4.Delta {
		return hw4.ErrNotEnoughFuel
	}
	cf.obj.setNeededFuel(neededFuel)

	return nil
}

func (cf *CheckFuel) getDistance(start, finish *pkg.Vector) float64 {
	dx := finish.X - start.X
	dy := finish.Y - start.Y
	return math.Sqrt(dx*dx + dy*dy)
}

func (cf *CheckFuel) calculateFuel(distance float64) float64 {
	return distance * 1 // todo расчет можно усложнить
}
