package test

import (
	"testing"
	"fracture/algorithm"
	"fracture"
)

func TestFracture(test *testing.T) {
	routine := algorithm.GetJuliaRoutine(-0.8 + 0.156i)
	fracture.GetImage(routine, 800, 600, 0.000000, 0.000000, 0.005000, "testX2.png")
		fracture.F()

//	is := view.NewImageSaver(800, 600, "testX.png")
//	is.Init()
//	routine := algorithm.GetJuliaRoutine(-0.8 + 0.156i)
//	routine(is, 800, 600, 0.003020, 0.001040, 0.000050)
//			is.FlushLoop()
}
