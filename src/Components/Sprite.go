package Components

import "github.com/hajimehoshi/ebiten/v2"

type ISprite interface {
	SetSprite(e *ebiten.Image)
	GetSprite() *ebiten.Image
	Animate()
	GetWidth() float64
	GetHeight() float64
}

type Sprite struct {
	currentSprite *ebiten.Image

	imageWidth  float64
	imageHeight float64
}

func (s *Sprite) GetWidth() float64 {
	return s.imageWidth
}

func (s *Sprite) GetHeight() float64 {
	return s.imageHeight
}

func (s *Sprite) SetSprite(e *ebiten.Image) {
	s.currentSprite = e

	w, h := s.currentSprite.Size()
	if int(s.imageWidth) != w {
		s.imageWidth = float64(w)
	}

	if int(s.imageHeight) != h {
		s.imageHeight = float64(h)
	}

}

func (s *Sprite) GetSprite() *ebiten.Image {
	return s.currentSprite
}

func (s *Sprite) Animate() {
	//TODO implement me
	panic("implement me")
}
