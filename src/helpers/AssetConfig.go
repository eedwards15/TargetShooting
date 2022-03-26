package helpers

import (
	"TargetShooting/assets"
	"TargetShooting/src/models"
	"encoding/json"
	"fmt"
	"log"
	"path"
)

func AssetConfigHelper() ([]*models.AssetConfig, error) {
	files, err := assets.AssetsFileSystem.ReadDir("configs")
	if err != nil {
		fmt.Println("Failed Loading Configs")
		log.Fatal(err)
	}

	assetConfigs := []*models.AssetConfig{}
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fileValue, _ := assets.AssetsFileSystem.ReadFile(path.Join("configs", file.Name()))
		assetConfig := models.AssetConfig{}

		json.Unmarshal(fileValue, &assetConfig)
		assetConfigs = append(assetConfigs, &assetConfig)
	}

	return assetConfigs, nil
}
