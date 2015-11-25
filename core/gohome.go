package core

import (
	"github.com/kkamperschroer/GoHome/config"
	"github.com/kkamperschroer/GoHome/models"
	"github.com/kkamperschroer/GoHome/plugin_loader"
	"github.com/pivotal-golang/lager"
)

type GoHome interface {
	Go()
}

type GoHomeInstance struct {
	logger  lager.Logger
	plugins []models.Plugin
}

func NewGoHome(config *config.Config, logger lager.Logger) (GoHome, error) {
	logger.Debug("Initializing GoHome")

	pluginLoader := plugin_loader.NewPluginLoader(config, logger)
	plugins, err := pluginLoader.LoadPlugins()

	if err != nil {
		logger.Error("Error loading plugins", err)
		return nil, err
	}

	logger.Debug("Validating configuration...")

	configValidator := config_validator.NewConfigValidator(config, plugins)
	err = configValidator.ValidateConfig()

	if err != nil {
		logger.Error("Invalid configuration", err)
		return nil, err
	}

	return &GoHomeInstance{
		logger:  logger,
		plugins: plugins,
	}, nil
}

func (g *GoHomeInstance) Go() {
	g.logger.Debug("GoHome started!")
}
