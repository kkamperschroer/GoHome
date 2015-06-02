package config

import (
	"github.com/kkamperschroer/GoHome/models"
)

func LoadPluginManifest(manifestLocation string) (*models.Plugin, error) {
	var plugin models.Plugin
	err := loadConfigIntoStruct(manifestLocation, &plugin)
	if err != nil {
		return nil, err
	}

	return &plugin, nil
}
