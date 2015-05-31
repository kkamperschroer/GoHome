package core

import (
	"github.com/kkamperschroer/GoHome/config"
	"github.com/pivotal-golang/lager"
)

type GoHome interface {
	Go()
}

type GoHomeInstance struct {
	logger lager.Logger
}

func NewGoHome(config *config.Config, logger lager.Logger) (GoHome, error) {
	logger.Debug("Initializing GoHome")

	return &GoHomeInstance{
		logger: logger,
	}, nil
}

func (g *GoHomeInstance) Go() {
	g.logger.Debug("GoHome started!")
}
