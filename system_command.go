package main

import(
	"fmt";
	)

type SystemCommand struct {
	Event // By doing thing, essentially inheriting from event

}

func (systemCommand* SystemCommand) Init() {
	systemCommand.name = "SystemCommand"

	fmt.Println("SystemCommand initialized.");
}

