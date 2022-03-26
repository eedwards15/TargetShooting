package scenes

import (
	"TargetShooting/src/helpers"
	"TargetShooting/src/systems"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/keyboard/keyboard"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image/color"
)

const (
	tileSize      = 32
	titleFontSize = fontSize * 1.5
	fontSize      = 36
)

type MainMenu struct {
	titleArcadeFont font.Face
	keys            []ebiten.Key
}

func NewMainMenu() *MainMenu {
	m := &MainMenu{}
	return m
}

func (m MainMenu) GetName() string {
	return "Main Menu"
}

func (m *MainMenu) Init() {
	fontFromSystem := systems.ASSETSYSTEM.Assets["Global"].Fonts["KenneyPixel"]
	systems.MUSICSYSTEM.PlaySong()
	const dpi = 72
	m.titleArcadeFont, _ = opentype.NewFace(fontFromSystem, &opentype.FaceOptions{
		Size:    titleFontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
}

func (m *MainMenu) Draw(screen *ebiten.Image) {
	text.Draw(screen, "Go TargetShooting", m.titleArcadeFont, helpers.CenterTextXPos("Go TargetShooting", m.titleArcadeFont, systems.WINDOWMANAGER.ScreenWidth), (systems.WINDOWMANAGER.ScreenHeight / 2), color.White)
}

func (m *MainMenu) Update() error {
	m.keys = inpututil.AppendPressedKeys(m.keys[:0])

	for _, p := range m.keys {
		_, ok := keyboard.KeyRect(p)

		if p.String() == "Enter" {
			systems.SCENEMANAGER.Push(&MainScene{})
		}

		if !ok {
			continue
		}

	}

	return nil
}
