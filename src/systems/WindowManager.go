package systems

import "log"

var (
	WINDOWMANAGER *WindowManager
)

type WindowManager struct {
	ScreenWidth  int
	ScreenHeight int
}

func InitWindowManager(width int, height int) {
	if WINDOWMANAGER == nil {
		if width == 0 {
			log.Fatalf("Missing Width")
		}
		if height == 0 {
			log.Fatalf("Missing Height")
		}

		windowManger := &WindowManager{
			ScreenWidth:  width,
			ScreenHeight: height,
		}
		WINDOWMANAGER = windowManger
	}

}

func (windowManagerClass WindowManager) Center() (float64, float64) {
	return float64(windowManagerClass.ScreenWidth / 2), float64(windowManagerClass.ScreenHeight / 2)
}
