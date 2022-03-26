package src

import (
	"TargetShooting/src/scenes"
	"TargetShooting/src/systems"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	keys []ebiten.Key
}

func (gameClass *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return systems.WINDOWMANAGER.ScreenWidth, systems.WINDOWMANAGER.ScreenHeight
}

func NewGame() *Game {
	g := &Game{}
	g.init()
	return g
}

func (gameClass *Game) init() {
	systems.SCENEMANAGER.Push(scenes.NewMainMenu())
}

func (gameClass *Game) Update() error {
	systems.SCENEMANAGER.CurrentScene.Update()
	return nil
}

func (gameClass *Game) Draw(screen *ebiten.Image) {
	systems.SCENEMANAGER.CurrentScene.Draw(screen)
}
