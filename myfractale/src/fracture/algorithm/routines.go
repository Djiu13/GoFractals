package algorithm

import (
	"fmt"
	"fracture/data"
	"math"
)

func MandelbrotRoutine(v View, width, height int, ox, oy, precision float64) {
	matrix := data.NewMatrix(width, height)
	channel := make(chan data.Pair, 100)

	n := math.Log(precision/0.005) / math.Log(0.9)
	iter := int(300 + n*5)

	v.Log(fmt.Sprintf("iter: %d", iter))

	algo := GetMandelbrot(iter)
	treatment := PixelTreatment(ox, oy, precision, matrix, channel, algo)

	go v.Listen(channel, matrix)

	go func() {
		data.LinearCourse(matrix.Width(), matrix.Height())(treatment)
		close(channel)
		v.Log("done")
	}()
}

func GetJuliaRoutine(c complex128) func(v View, width, height int, ox, oy, precision float64) {
	return func(v View, width, height int, ox, oy, precision float64) {
		matrix := data.NewMatrix(width, height)
		channel := make(chan data.Pair, 100)

		n := math.Log(precision/0.005) / math.Log(0.9)
		iter := int(300 + n*5)

		v.Log(fmt.Sprintf("iter: %d", iter))

		algo := GetJulia(iter, c)
		treatment := PixelTreatment(ox, oy, precision, matrix, channel, algo)

		go v.Listen(channel, matrix)

		go func() {
			data.LinearCourse(matrix.Width(), matrix.Height())(treatment)
			close(channel)
			v.Log("done")
		}()
	}
}
