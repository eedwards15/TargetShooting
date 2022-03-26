package Entities

import (
	"TargetShooting/src/Components"
	"TargetShooting/src/helpers"
	"TargetShooting/src/models"
	"github.com/hajimehoshi/ebiten/v2"
)

type ITarget interface {
	GetTimeToAdd() int
	Getvalue() int
	Draw(screen *ebiten.Image)
	Update()

	LifeSpanTicker()
	IsDead() bool

	GetTransform() Components.Transform
	TargetHit(mX, mY float64) bool
}

type Target struct {
	targetValue int
	timeValue   int

	Components.Transform
	Components.Sprite

	targetLifeSpan int

	isDead bool
}

func (t *Target) TargetHit(mX, mY float64) bool {
	if helpers.BoxCollision(mX, mY, 1, 1, t.XPos, t.YPos, t.GetWidth(), t.GetHeight()) {
		t.isDead = true
		return true
	}
	return false
}

func (t Target) GetTransform() Components.Transform {
	return t.Transform
}

func (t Target) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(t.XPos, t.YPos)
	screen.DrawImage(t.GetSprite(), op)
}

func (t *Target) IsDead() bool {
	return t.isDead
}

func (t *Target) LifeSpanTicker() {
	t.isDead = true
}

func (t Target) Update() {}

func NewTarget(transform models.Vector, sprite *ebiten.Image, targetValue int, timeValue int) *Target {
	target := &Target{}
	target.Transform.Vector = transform
	target.SetSprite(sprite)
	target.targetValue = targetValue
	target.timeValue = timeValue
	target.targetLifeSpan = 2
	target.isDead = false
	return target
}

func (t Target) GetTimeToAdd() int {
	return t.timeValue
}

func (t Target) Getvalue() int {
	return t.targetValue
}
