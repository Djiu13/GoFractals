// Interface

package fracture

import (
	"fracture/algorithm"
	"fracture/view"
)

// Comment
func F() string {
	return "bonjour"
}

func GetImage(routine algorithm.Routine, width int, height int, x float64, y float64, p float64,
	filePath string) {
	is := view.NewImageSaver(width, height, filePath)
	is.Init()
	routine(is, width, height, x, y, p)
	is.FlushLoop()
}
