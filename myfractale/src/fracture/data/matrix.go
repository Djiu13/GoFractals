package data

// A simple two-dimension matrix implementation
type Matrix struct {
	matrix []float64
	width  int
	height int
}

// Constructor of Matrix structure
func NewMatrix(width int, height int) (res *Matrix) {
	res = new(Matrix)
	res.matrix = make([]float64, width*height)
	res.width, res.height = width, height

	return
}

func (mtrx *Matrix) At(i int, j int) (res float64) {
	if i < mtrx.width && j < mtrx.height {
		res = mtrx.matrix[j*mtrx.width+i]
	}

	return
}

func (mtrx *Matrix) Set(i int, j int, value float64) {
	if i < mtrx.width && j < mtrx.height {
		mtrx.matrix[j*mtrx.width+i] = value
	}
}

func (mtrx *Matrix) SetInt(i, j, value int) {
	mtrx.Set(i, j, float64(value))
}

func (mtrx *Matrix) Width() int {
	return mtrx.width
}

func (mtrx *Matrix) Height() int {
	return mtrx.height
}

// A Treatment is a function which is apply on each point during the matrix course
type Treatment func(int, int)

// A MatrixCourse is a way to travel into a matrix
type MatrixCourse func(Treatment)

func LinearCourse(width, height int) func(treatment Treatment) {
	return func(treatment Treatment) {
		for j := 0; j < height; j++ {
			for i := 0; i < width; i++ {
				treatment(i, j)
			}
		}
	}
}
