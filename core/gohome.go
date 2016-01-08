package core

import (
	"github.com/kkamperschroer/GoHome/config"
	"github.com/kkamperschroer/GoHome/config_validator"
	"github.com/kkamperschroer/GoHome/models"
	"github.com/kkamperschroer/GoHome/plugin_loader"
	"github.com/kkamperschroer/GoHome/plugin_server"
	"github.com/kkamperschroer/GoHome/rules_engine"
	"github.com/pivotal-golang/lager"
	"time"
)

type GoHome interface {
	Go() error
}

type goHome struct {
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

	pluginServer := plugin_server.NewPluginServer(logger)
	err = pluginServer.Start()

	if err != nil {
		logger.Error("Error starting up plugin server", err)
		return nil, err
	}

	rulesEngine, err := rules_engine.NewRulesEngine(config, logger, plugins)
	if err != nil {
		logger.Error("Error creating rules engine", err)
		return nil, err
	}

	err = rulesEngine.Start()
	if err != nil {
		logger.Error("Error starting rules engine", err)
		return nil, err
	}

	return &goHome{
		logger:  logger,
		plugins: plugins,
	}, nil
}

func (g *goHome) Go() error {
	g.logger.Debug("GoHome started! Waiting for events to fire...")

	for {
		// TODO -- how can i just wait for goroutines? channels, dummy
		time.Sleep(2 * time.Minute)
		g.logger.Info("GoHome sleeping...")
	}

	return nil
}
