package main

import(
	"fmt";
	"time";
	)

type SystemCommand struct {
	name string
	lastFired time.Time
	pollingInterval time.Time
}

func (systemCommand* SystemCommand) Init() {
	systemCommand.name = "SystemCommand"

	fmt.Println("SystemCommand initialized.");
}

func (s* SystemCommand) GetName() string {
	return s.name
}

func (s* SystemCommand) GetLastFired() time.Time {
	return s.lastFired
}

func (s* SystemCommand) Fire() {
	s.lastFired = time.Now()

	fmt.Println("Firing system command!")
}
