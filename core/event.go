package core

import (
	"time"
)

type Event interface {
	LastFired() time.Time
}
