package burnfuel

type BurnFuel struct {
	obj BurnableFuel
}

func NewBurnFuel(obj BurnableFuel) *BurnFuel {
	return &BurnFuel{obj}
}

func (bf *BurnFuel) Execute() error {
	neededFuel, err := bf.obj.getNeededFuel()
	if err != nil {
		return err
	}
	bf.obj.setFuel(neededFuel)

	return nil
}
