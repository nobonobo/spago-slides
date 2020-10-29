package slide

import (
	"github.com/nobonobo/spago"
)

// Render ...
func (c *Slides) Render() spago.HTML {
	return spago.Tag("body", 
		spago.C(c.Slides...),
		spago.Tag("div", 			
			spago.A("class", spago.S(`control`)),
			spago.Tag("button", 				
				spago.A("class", spago.S(`btn`)),
				spago.Event("click", c.OnPrev),
				spago.T(`◀️`),
			),
			spago.Tag("button", 				
				spago.A("class", spago.S(`btn`)),
				spago.Event("click", c.OnNext),
				spago.T(`▶️`),
			),
		),
	)
}
