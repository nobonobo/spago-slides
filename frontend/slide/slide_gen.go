package slide

import (
	"github.com/nobonobo/spago"
)

// Render ...
func (c *Slide) Render() spago.HTML {
	return spago.Tag("div", 		
		spago.A("class", spago.S(`slide`)),
		spago.A("id", spago.S(``, spago.S(c.ID), ``)),
		spago.Tag("div", 			
			spago.A("class", spago.S(`card`)),
			spago.Tag("div", 				
				spago.A("class", spago.S(`card-body`)),
				spago.UnsafeHTML(c.Content),
			),
		),
	)
}
