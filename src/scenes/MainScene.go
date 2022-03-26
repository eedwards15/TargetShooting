package scenes

import (
	"TargetShooting/src/Entities"
	"TargetShooting/src/helpers"
	"TargetShooting/src/systems"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image/color"
	"strconv"
	"time"
)

type MainScene struct {
	spawner *systems.Spawner
	Score   int
	targets []Entities.ITarget

	targetTimer     time.Timer
	titleArcadeFont font.Face

	gameTimer time.Time
	playerFx  *audio.Player

	bullets int

	difficulty  int
	endingScore font.Face
	gameOver    bool
}

//Should be the Config Name.
func (m *MainScene) GetName() string { return "Main Scene" }

func (m *MainScene) Init() {
	m.spawner = systems.NewSpawnerSystem()
	m.Score = 0
	m.difficulty = 0
	m.targets = []Entities.ITarget{}
	m.gameTimer = time.Now().Add(time.Duration(5) * time.Second)
	const dpi = 72
	fontFromSystem := systems.ASSETSYSTEM.Assets["Global"].Fonts["KenneyPixel"]
	m.titleArcadeFont, _ = opentype.NewFace(fontFromSystem, &opentype.FaceOptions{
		Size:    60,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	m.endingScore, _ = opentype.NewFace(fontFromSystem, &opentype.FaceOptions{
		Size:    titleFontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	soundEffect := systems.ASSETSYSTEM.Assets["Global"].SoundEffects["Shot"]
	m.playerFx, _ = audio.CurrentContext().NewPlayer(soundEffect)

	m.bullets = 3

}

func (m MainScene) Draw(screen *ebiten.Image) {
	backgroundOP := &ebiten.DrawImageOptions{}
	backgroundOP.GeoM.Scale(2, 2)
	screen.DrawImage(systems.ASSETSYSTEM.Assets["Global"].Images["Background"], backgroundOP)

	if m.gameOver {
		text.Draw(screen, "Game Over", m.endingScore, helpers.CenterTextXPos("Game Over", m.endingScore, systems.WINDOWMANAGER.ScreenWidth), (systems.WINDOWMANAGER.ScreenHeight/2)-200, color.White)
		text.Draw(screen, "Score: "+strconv.Itoa(m.Score), m.endingScore, helpers.CenterTextXPos("Score: "+strconv.Itoa(m.Score), m.endingScore, systems.WINDOWMANAGER.ScreenWidth), (systems.WINDOWMANAGER.ScreenHeight/2)-100, color.White)
		text.Draw(screen, "Press Enter to play again: ", m.endingScore, helpers.CenterTextXPos("Press Enter to play again", m.endingScore, systems.WINDOWMANAGER.ScreenWidth), systems.WINDOWMANAGER.ScreenHeight/2, color.White)

		return
	}

	for i := 0; i < len(m.targets); i++ {
		if !m.targets[i].IsDead() {
			m.targets[i].Draw(screen)
		}
	}

	for i := 0; i < m.bullets; i++ {
		backgroundOP := &ebiten.DrawImageOptions{}
		backgroundOP.GeoM.Translate(float64(50+(i*28)), float64(systems.WINDOWMANAGER.ScreenHeight-70))
		screen.DrawImage(systems.ASSETSYSTEM.Assets["Global"].Images["Bullet"], backgroundOP)
	}

	text.Draw(screen, "Score: "+strconv.Itoa(m.Score), m.titleArcadeFont, 20, 30, color.White)
	text.Draw(screen, "Seconds Remaining: "+strconv.Itoa(int(m.gameTimer.Sub(time.Now()).Seconds())), m.titleArcadeFont, 20, 60, color.White)
}

func (m *MainScene) Reset() {
	m.Score = 0
	m.difficulty = 0
	m.targets = []Entities.ITarget{}
	m.gameTimer = time.Now().Add(time.Duration(5) * time.Second)
	m.gameOver = false
}

func (m *MainScene) Update() error {

	if int(m.gameTimer.Sub(time.Now()).Seconds()) <= 0 {
		m.gameOver = true

		if inpututil.IsKeyJustReleased(ebiten.KeyEnter) {
			m.Reset()
		}
		return nil
	}

	newTarget := m.spawner.SpawnNewEnemy()
	if newTarget != nil {
		if m.Score > 500 && m.difficulty < 1300 {
			m.difficulty += 100

		}

		time.AfterFunc(time.Duration(2000-m.difficulty)*time.Millisecond, newTarget.LifeSpanTicker)
		m.targets = append(m.targets, newTarget)
	}

	for i := 0; i < len(m.targets); i++ {
		m.targets[i].Update()
	}

	temp := []Entities.ITarget{}
	for i := 0; i < len(m.targets); i++ {
		if !m.targets[i].IsDead() {
			temp = append(temp, m.targets[i])
		}
	}
	m.targets = temp

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) && m.bullets > 0 {
		m.playerFx.Rewind()
		m.playerFx.Play()

		m.bullets -= 1

		mX, mY := ebiten.CursorPosition()

		for i := 0; i < len(m.targets); i++ {

			if m.targets[i].TargetHit(float64(mX), float64(mY)) {
				m.Score += m.targets[i].Getvalue()
				m.gameTimer = m.gameTimer.Add(time.Duration(m.targets[i].GetTimeToAdd()) * time.Second)
			}
		}
	}

	if inpututil.IsKeyJustReleased(ebiten.KeyR) {
		m.bullets = 3
	}

	return nil
}
