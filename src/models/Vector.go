package models

type Vector struct {
	XPos float64
	YPos float64
}

func NewVector(x, y float64) *Vector {
	v := &Vector{
		XPos: x,
		YPos: y,
	}

	return v
}