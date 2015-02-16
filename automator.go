package main

import(
	"fmt";
	"container/list";
	)

type Automator struct {
	events list.List
	actions list.List
}

func (a *Automator) Init() {
	fmt.Println("In automator init")
	a.events.Init()
	a.actions.Init()

	// TODO -- Read in the events and actions from some config

	a.buildEvents();
}

func (a *Automator) buildEvents() {
	a.events.PushBack("test")
}

func (a *Automator) Run() {
	fmt.Println("In automator run")
}
