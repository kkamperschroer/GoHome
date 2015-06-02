package plugin_loader

import (
	"fmt"
	"github.com/kkamperschroer/GoHome/config"
	"github.com/kkamperschroer/GoHome/models"
	"github.com/pivotal-golang/lager"
)

type PluginLoader interface {
	LoadPlugins() ([]models.Plugin, error)
}

type PluginLoaderInstance struct {
	config *config.Config
	logger lager.Logger
}

func NewPluginLoader(config *config.Config, logger lager.Logger) PluginLoader {
	return &PluginLoaderInstance{
		config: config,
		logger: logger,
	}
}

func (p *PluginLoaderInstance) LoadPlugins() ([]models.Plugin, error) {
	returnPlugins := make([]models.Plugin, len(p.config.Plugins))

	for _, pluginConfig := range p.config.Plugins {
		p.logger.Info(fmt.Sprintf("Loading plugin '%s'", pluginConfig.Name))

		plugin, err := p.loadPlugin(pluginConfig.ManifestLocation)
		if err != nil {
			p.logger.Error(fmt.Sprintf("Error loading plugin '%s'", pluginConfig.Name), err)
			return nil, err
		}

		returnPlugins = append(returnPlugins, *plugin)
	}

	p.logger.Info("Plugins loaded successfully")
	return returnPlugins, nil
}

func (p *PluginLoaderInstance) loadPlugin(manifestLocation string) (*models.Plugin, error) {
	plugin, err := config.LoadPluginManifest(manifestLocation)
	if err != nil {
		return nil, err
	}

	return plugin, nil
}
