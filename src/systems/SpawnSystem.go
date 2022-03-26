package systems

import (
	"TargetShooting/assets"
	"TargetShooting/src/Entities"
	"TargetShooting/src/models"
	"encoding/json"
	"io/fs"
	"math/rand"
	"path"
	"time"
)

type Spawner struct {
	enemyConfigs    []models.TargetConfig
	LAST_SPAWN_TIME time.Time
	coolDown        float64
}

func NewSpawnerSystem() *Spawner {
	eS := &Spawner{}
	eS.LAST_SPAWN_TIME = time.Now().Add(-time.Duration(2) * time.Second)
	eS.coolDown = 2
	configs, _ := assets.AssetsFileSystem.ReadDir("settings/targets")
	for i := 0; i < len(configs); i++ {
		enemyConfig := unPackConfig(configs, i)
		eS.enemyConfigs = append(eS.enemyConfigs, enemyConfig)
	}
	return eS
}

func unPackConfig(configs []fs.DirEntry, i int) models.TargetConfig {
	fileValue, _ := assets.AssetsFileSystem.ReadFile(path.Join("settings/targets", configs[i].Name()))
	enemyConfig := models.TargetConfig{}
	json.Unmarshal(fileValue, &enemyConfig)
	return enemyConfig
}

func (spawner *Spawner) SpawnNewEnemy() Entities.ITarget {

	if time.Now().Sub(spawner.LAST_SPAWN_TIME).Seconds() >= spawner.coolDown {
		spawner.LAST_SPAWN_TIME = time.Now()

		//Location To Render
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		xPos := float64(r1.Intn(WINDOWMANAGER.ScreenWidth - (100)))
		yPos := float64(r1.Intn(WINDOWMANAGER.ScreenHeight - 100))

		//Random Config
		s2 := rand.NewSource(time.Now().UnixNano())
		r2 := rand.New(s2)
		configIndex := r2.Intn(len(spawner.enemyConfigs))

		//Create Target
		targetConfig := spawner.enemyConfigs[configIndex]
		transform := models.NewVector(xPos, yPos)
		newTarget := Entities.NewTarget(*transform, ASSETSYSTEM.Assets["Global"].Images[targetConfig.Key], targetConfig.Value, targetConfig.TimeValue)

		return newTarget
	}

	return nil
}
