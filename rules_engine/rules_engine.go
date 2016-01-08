package rules_engine

import (
	"fmt"

	"github.com/kkamperschroer/GoHome/config"
	"github.com/kkamperschroer/GoHome/models"
	"github.com/pivotal-golang/lager"
)

type RulesEngine interface {
	Start() error
}

type rulesEngine struct {
	logger lager.Logger
	rules  []models.Rule
}

// Kyle -- this is where you left off. Creation of the rule engine. Currently just logging some stuff.
// Need to actually build the state machine for each rule.

func NewRulesEngine(config *config.Config, logger lager.Logger, plugins []models.Plugin) (RulesEngine, error) {
	for _, rule := range config.Rules {
		logger.Debug(fmt.Sprintf("Creating rule '%s'", rule.Name))
		for _, action := range rule.Actions {
			logger.Debug(fmt.Sprintf("Rule '%s': Creating action to '%s'", rule.Name, action.Description))
		}

		for _, condition := range rule.Conditions {
			logger.Debug(fmt.Sprintf("Rule '%s': Creating condition to check '%s'", rule.Name, condition.Description))
		}

		logger.Debug(fmt.Sprintf("Rule '%s': Creating event '%s'", rule.Name, rule.Event.Description))
	}

	return &rulesEngine{
		logger: logger,
	}, nil
}

func (r *rulesEngine) Start() error {
	r.logger.Debug("Rules engine started. Beginning event loops")

	return nil
}
