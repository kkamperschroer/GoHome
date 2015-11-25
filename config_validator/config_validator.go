package config_validator

import (
	"github.com/kkamperschroer/GoHome/config"
	"github.com/kkamperschroer/GoHome/models"
)

type ConfigValidator interface {
	ValidateConfig() error
}

type configValidator struct {
	config  *config.Config
	plugins []models.Plugin
}

func NewConfigValidator(config *config.Config, plugins []models.Plugin) ConfigValidator {
	return &configValidator{
		config:  config,
		plugins: plugins,
	}
}

func (c *configValidator) ValidateConfig() error {
	// TODO
	return nil
}
