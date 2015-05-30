package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Rules []RuleConfig `json:"rules"`
}

type RuleConfig struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Event       EventConfig       `json:"event"`
	Conditions  []ConditionConfig `json:"conditions"`
	Actions     []ActionConfig    `json:"actions"`
}

type EventConfig struct {
	Description     string `json:"description"`
	PluginId        string `json:"plugin_id"`
	PluginEventName string `json:"plugin_event_name"`
	Config          interface{}
}

type ConditionConfig struct {
	Description         string `json:"description"`
	IsAnd               bool   `json:"isAnd"`
	IsOr                bool   `json:"isOr"`
	IsNot               bool   `json:"isNot"`
	PluginId            string `json:"plugin_id"`
	PluginConditionName string `json:"plugin_condition_name"`
	Config              interface{}
}

type ActionConfig struct {
	Description      string `json:"description"`
	PluginId         string `json:"plugin_id"`
	PluginActionName string `json:"plugin_action_name"`
	Config           interface{}
}

func LoadConfigData(filepath string) (*Config, error) {
	configBytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var configData Config
	err = json.Unmarshal(configBytes, &configData)
	if err != nil {
		return nil, err
	}

	return &configData, nil
}
