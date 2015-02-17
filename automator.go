package main

import(
	"fmt";
	"container/list";
	)

type Automator struct {
	events list.List
	actions list.List
}

func (automator *Automator) Init() {
	fmt.Println("In automator init")
	automator.events.Init()
	automator.actions.Init()

	automator.BuildEvents();
}

func (automator *Automator) BuildEvents() {
	// TODO -- Read in the events and actions from some config
	
	sysCmd := new(SystemCommand)
	sysCmd.Init()
	automator.events.PushBack(sysCmd)
}

func (automator *Automator) Run() {
	fmt.Println("In automator run")

}


