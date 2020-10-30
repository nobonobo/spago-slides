package actions

import "github.com/nobonobo/spago/dispatcher"

const (
	// Refresh ...
	Refresh dispatcher.Action = iota + 1
	// NextStep ...
	NextStep
	// PrevStep ...
	PrevStep
	// ReLoad ...
	ReLoad
)
