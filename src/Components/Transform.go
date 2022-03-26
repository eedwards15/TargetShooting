package Components

import "TargetShooting/src/models"

type ITransform interface {
	GetXPos() float64
	GetYPos() float64
	Move(x, y float64)
}

type Transform struct {
	models.Vector
}

func (t Transform) GetXPos() float64 {
	return t.XPos
}

func (t Transform) GetYPos() float64 {
	return t.YPos
}

func (t *Transform) Move(x, y float64) {
	t.YPos += y
	t.XPos += x
}
