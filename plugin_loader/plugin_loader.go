package plugin_loader

import (
	"fmt"
	"github.com/kkamperschroer/GoHome/config"
	"github.com/kkamperschroer/GoHome/models"
	"github.com/pivotal-golang/lager"
	"io/ioutil"
	"os"
	"path/filepath"
)

type PluginLoader interface {
	LoadPlugins() ([]models.Plugin, error)
}

type pluginLoader struct {
	config *config.Config
	logger lager.Logger
}

func NewPluginLoader(config *config.Config, logger lager.Logger) PluginLoader {
	return &pluginLoader{
		config: config,
		logger: logger,
	}
}

func (p *pluginLoader) LoadPlugins() ([]models.Plugin, error) {
	fileInfos, err := ioutil.ReadDir(p.config.Plugins.Location)
	if err != nil {
		return nil, err
	}

	returnPlugins := make([]models.Plugin, len(fileInfos))

	for _, file := range fileInfos {
		if file.IsDir() {
			possiblePluginDir := filepath.Join(p.config.Plugins.Location, file.Name())
			p.logger.Debug("Looking in '" + possiblePluginDir + "' for plugin")

			plugin, err := p.loadPlugin(possiblePluginDir)
			if err != nil {
				if os.IsNotExist(err) {
					p.logger.Info(fmt.Sprintf("Ignoring folder at '%s' due to lack of manifest.json", possiblePluginDir))
				} else {
					p.logger.Error(fmt.Sprintf("Error loading manifest at '%s'. Skipping", possiblePluginDir), err)
				}
			} else {
				p.logger.Debug(fmt.Sprintf("Loaded a plugin called '%s' from '%s'", plugin.Name, possiblePluginDir))
				returnPlugins = append(returnPlugins, *plugin)
			}
		}
	}

	p.logger.Info("Plugins loaded successfully")
	return returnPlugins, nil
}

func (p *pluginLoader) loadPlugin(dirPath string) (*models.Plugin, error) {
	manifestLocation := filepath.Join(dirPath, "manifest.json")

	plugin, err := config.LoadPluginManifest(manifestLocation)
	if err != nil {
		return nil, err
	}

	return plugin, nil
}
