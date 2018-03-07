package algorithm

import (
	"math/cmplx"
)

func GetMandelbrot(itermax int) func(complex128) float64 {
	return func(c complex128) float64 {
		iter := 0
		var zn, zn1 complex128

		for ; iter < itermax; iter++ {
			zn1 = zn*zn + c

			if cmplx.Abs(zn1) > 2 {
				break
			}

			zn = zn1
		}

		if iter == itermax {
			iter = -1
		}

		return float64(iter)
	}
}

func GetJulia(itermax int, c complex128) func(complex128) float64 {
	return func(zn complex128) float64 {
		iter := 0
		var zn1 complex128

		for ; iter < itermax; iter++ {
			zn1 = zn*zn + c

			if cmplx.Abs(zn1) > 2 {
				break
			}

			zn = zn1
		}

		return float64(iter)
	}
}
