package slide

import (
	"github.com/nobonobo/spago"
)

//go:generate spago generate -c Slide -p slide slide.html

// Slide  ...
type Slide struct {
	spago.Core
	ID      string
	Content string
}
