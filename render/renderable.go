// Package render provides several types of renderable entities,
// methods for loading images from file, and methods for drawing
// those entities to screen.
package render

import (
	"golang.org/x/exp/shiny/screen"
	"image"
)

// A Renderable is anything which can
// be drawn at a given draw layer, undrawn,
// and set in a particular position.
//
// Basic Implementing struct: Sprite
type Renderable interface {
	// As the engine currently exists,
	// the buffer which is passed into draw
	// is always the same. This leads to
	// several parts of the engine being
	// reliant on shiny/screen when they
	// could call out to somewhere else to
	// determine what they are drawn onto.
	//
	// On the other hand, this allows manually
	// duplicating renderables onto multiple
	// buffers, but in certain implementations
	// (i.e Animation) would have unintended
	// consequences.
	Draw(buff screen.Buffer)
	GetRGBA() *image.RGBA
	// Basic Implementing struct: Point
	ShiftX(x float64)
	ShiftY(y float64)
	SetPos(x, y float64)
	// Basic Implementing struct: Layered
	GetLayer() int
	SetLayer(l int)
	UnDraw()
}
