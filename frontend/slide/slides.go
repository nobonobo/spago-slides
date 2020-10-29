package slide

import (
	"syscall/js"

	"github.com/nobonobo/spago"
	"github.com/nobonobo/spago-slides/frontend/actions"
	"github.com/nobonobo/spago/dispatcher"
)

//go:generate spago generate -c Slides -p slide slides.html

var (
	location = js.Global().Get("location")
)

// Slides  ...
type Slides struct {
	spago.Core
	Slides []spago.Component
}

// OnPrev ...
func (c *Slides) OnPrev(ev js.Value) {
	dispatcher.Dispatch(actions.PrevStep)
}

// OnNext ...
func (c *Slides) OnNext(ev js.Value) {
	dispatcher.Dispatch(actions.NextStep)
}
