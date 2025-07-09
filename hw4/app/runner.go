package app

import (
	"architecture_otus/hw4"
	"architecture_otus/hw4/burnfuel"
	"architecture_otus/hw4/changevelocity"
	"architecture_otus/hw4/checkfuel"
	"architecture_otus/hw4/move"
	"architecture_otus/hw4/rotate"
)

func Move(sp hw4.SpaceObject) error {
	queue := []hw4.Command{
		checkfuel.NewCheckFuel(checkfuel.NewCheckableFuelAdapter(sp)),
		move.NewMove(move.NewMovableObjectAdapter(sp)),
		burnfuel.NewBurnFuel(burnfuel.NewBurnableFuelAdapter(sp)),
	}

	for _, cmd := range queue {
		err := cmd.Execute()
		if err != nil {
			return err
		}
	}
	return nil
}

func Rotate(sp hw4.SpaceObject) error {
	queue := []hw4.Command{
		changevelocity.NewChangeVelocity(changevelocity.NewChangeableVelocityAdapter(sp)),
		rotate.NewRotate(rotate.NewRotatableObjectAdapter(sp)),
	}

	for _, cmd := range queue {
		err := cmd.Execute()
		if err != nil {
			return err
		}
	}
	return nil
}
