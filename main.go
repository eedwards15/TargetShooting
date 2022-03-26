package main

import (
	"TargetShooting/src"
	"TargetShooting/src/systems"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

func main() {
	systems.InitAssetSystem()
	systems.InitSceneManager()
	systems.InitWindowManager(1920, 1080)
	systems.InitMusicSystem(systems.ASSETSYSTEM.Assets["Global"].BackgroundMusic)

	g := src.NewGame()
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowTitle("Go TargetShooting")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
