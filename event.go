package main

import(
	"fmt";
	"time";
	)

type Event struct {
     name string
     last_fire time.Time
}

func (event* Event) Fire() {
	event.last_fire = time.Now()

	fmt.Println("Event fired! %+v", event)
}
