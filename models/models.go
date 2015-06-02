package models

type Plugin struct {
	Name       string            `json:"name"`
	Version    string            `json:"version"`
	Events     []PluginEvent     `json:"events"`
	Conditions []PluginCondition `json:"conditions"`
	Actions    []PluginAction    `json:"actions"`
}

type PluginEvent struct {
	PluginComponent
}

type PluginCondition struct {
	PluginComponent
}

type PluginAction struct {
	PluginComponent
}

type PluginComponent struct {
	Name  string `json:"name"`
	Entry string `json:"entry"`
}
