package algorithm

import (
	"fracture/data"
)

// Each pixel is computed independently
func PixelTreatment(ox, oy, precision float64, matrix *data.Matrix, channel chan data.Pair, algo func(complex128) float64) func(int, int) {
	return func(i, j int) {
		x := ox - float64(matrix.Width())*precision/2.0 + float64(i)*precision
		y := oy - float64(matrix.Height())*precision/2.0 + float64(j)*precision

		value := algo(complex(x, y))

		matrix.Set(i, j, value)

		channel <- data.Pair{i, j}
	}
}

// We keep the trace of the complex in the grid and increment the related value in the matrix
func TraceTreatment(ox, oy, precision float64, matrix *data.Matrix, channel chan data.Pair, algo func(complex128) []data.Pair) func(int, int) {
	return func(i, j int) {
		x := ox - float64(matrix.Width())*precision/2.0 + float64(i)*precision
		y := oy - float64(matrix.Height())*precision/2.0 + float64(j)*precision

		for _, v := range algo(complex(x, y)) {
			val := matrix.At(v.First, v.Second)

			matrix.Set(v.First, v.Second, val+1)
			channel <- data.Pair{v.First, v.Second}
		}
	}
}
