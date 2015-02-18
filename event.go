package main

import(
	"time";
	)

type Event interface {
	Init()
	GetName() string
	GetLastTimeFired() time.Time
	Fire()
}
