package algorithm

import "fracture/data"

type Routine func(v View, width, height int, ox, oy, precision float64)

// The view will deal the display of a fractal, during the computation, and receive the user input. Log will print the log in the screen (or whatever), LoopEvent will catch input user and Listen will get the modification of the computation with a chan
type View interface {
	Log(string)
	EventLoop(Routine)
	Listen(chan data.Pair, *data.Matrix)
	FlushLoop()
}
