# Oak 
### A pure Go game engine

----

## Installation
`go get -u github.com/OakmoundStudio/oak/...`

On linux, for audio, see [klangsynthese](https://github.com/200sc/klangsynthese) for audio installation requirements

## Usage
This is an example of the most basic oak program:
```
oak.AddScene("firstScene",
    func(prevScene string, inData interface{}) {}, // Initialization function
    func()bool{return true}, // Loop to continue or stop current scene
    func()(nextScene string, result *oak.SceneResult){return "firstScene", nil}) // Exit to transition to next scene
oak.Init("firstScene")
```
See the [examples](examples) folder for longer demos.

## Motiviation
The initial version of oak was made to support Oakmound Studio's game:
[Agent Blue](https://github.com/OakmoundStudio/AgentRelease) and was developed in parallel.
Oak supports Windows with no dependencies and Linux with limited audio dependencies.
 We hope that users will be able to make great pure Go games with oak and potentially improve oak.
 
 Because Oak wants to have as few dependencies as possible, Oak does not use OpenGL or [GLFW](https://github.com/go-gl/glfw).
 We're open to adding support for these in the future for performance gains, but we always want
 an alternative that requires zero or near-zero dependencies. (We are very sad about the linux audio 
 dependency and are considering writing an audio driver just to get rid of it.)

## Features
1. Window Rendering
    - Windows and key events through [shiny](https://github.com/golang/exp/tree/master/shiny)
    - Logical frame rate distinct from Draw rate
1. [Image Management](render)
    - `render.Renderable` interface
    - TileSheet Batch Loading
    - Manipulation
        - `render.Modifiable` interface
        - Built in Shaping, Coloring, Shading, ...
        - Some built ins via [gift](https://github.com/disintegration/gift)
        - extensible Modification syntax `func(image.Image) *image.RGBA`
        - Copying
    - Built in `Renderable` types
        - `Sprite`
        - Sheet `Animation`
        - `Sequence`, `Compound`, `Composite`
        - History-tracking `Reverting`
    - Primarily 2D
1. [Particle System](render/particle)
1. [Mouse Handling](mouse)
    - Click Collision
    - MouseEnter / MouseExit reaction events
    - Drag Handling
1. [Audio Support](audio)
    - From [klangsynthese](https://github.com/200sc/klangsynthese)
    - Batch Loading
    - Positional filters to pan and scale audio based on a listening position
1. [Collision](collision)
    - Collision R-Tree from [rtreego](https://github.com/dhconnelly/rtreego)
    - 2D Raycasting
    - Collision Spaces
        - Attachable to Objects
        - Auto React to collisions through events
        - OnHit bindings `func(s1,s2 *collision.Space)`
        - Start/Stop collision with targeted objects
1. [Physics System](physics)
    - Vectors
        - Attachable to Objects / Renderables
        - Momentum
        - Friction
        - Force / Pushing
1. [Event Handler, Bus](event)
    - PubSub system
    - `event.CID` can `Bind(fn,eventName)` and selectively `Trigger(eventName)` events
    - `GlobalBind` and `event.Trigger` for entity-independant 
1. [Timing utilities](timing)
    - Smoothed draw rate, frame rate tracking
    - FPS conversion to `time.Duration`
    - Manipulatable `time.Ticker` to readily change frame rate
1. [Shaping](shape)
    - Shapes from `func(x float64) (y float64)` equations
    - Shapes from `func(x,y, w, (h) int) bool` containment
    - Convert shapes into: 
        - Containment checks
        - Outlines
        - 2D arrays
1. [Custom Console Commands](debugConsole.go)
1. [Logging](dlog)
    - Controlled by config files
    - Filterable by string, debug level

## Package-specific Usage

... Pending! See examples or godoc!

## Contributions
See CONTRIBUTING.md