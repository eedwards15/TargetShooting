package systems

import (
	"TargetShooting/src/helpers"
	"TargetShooting/src/models"
	"sync"
)

var (
	ASSETSYSTEM *AssetSystem
)

type AssetSystem struct {
	Assets map[string]*models.SceneAssets
}

func InitAssetSystem() {
	ASSETSYSTEM = &AssetSystem{}
	ASSETSYSTEM.Assets = make(map[string]*models.SceneAssets)
	configValues, _ := helpers.AssetConfigHelper()

	var wg sync.WaitGroup
	for i := 0; i < len(configValues); i++ {
		wg.Add(1)
		r := configValues[i]
		go func(record *models.AssetConfig) {
			defer wg.Done()
			ASSETSYSTEM.Assets[record.Scene] = models.NewSceneAssets(record)

		}(r)
	}
	wg.Wait()

}
