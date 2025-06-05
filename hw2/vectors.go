package hw2

type Vector struct {
	x float64
	y float64
}

func (v *Vector) sum(a *Vector) *Vector {
	return &Vector{v.x + a.x, v.y + a.y}
}
