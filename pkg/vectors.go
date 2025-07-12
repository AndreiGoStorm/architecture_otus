package pkg

import "math"

type Vector struct {
	X float64
	Y float64
}

func (v *Vector) Sum(a *Vector) *Vector {
	return &Vector{v.X + a.X, v.Y + a.Y}
}

func (v *Vector) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
