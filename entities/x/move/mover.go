package move

import (
	"github.com/oakmound/oak/collision"
	"github.com/oakmound/oak/physics"
	"github.com/oakmound/oak/render"
)

// A Mover can move its position, renderable, and space. Unless otherwise documented,
// functions effecting a mover move all of its logical position, renderable, and space
// simultaneously.
type Mover interface {
	Vec() physics.Vector
	GetRenderable() render.Renderable
	GetDelta() physics.Vector
	GetSpace() *collision.Space
	GetSpeed() physics.Vector
}
