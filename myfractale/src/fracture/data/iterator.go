package data

// An Iterator is a tool helping to abstract the way we travel a Matrix. It does not provide a way to edit it.
type Iterator interface {
	// Position the iteration on the next point
	Next()
	// Verify if the iterator can do an other movement
	HasNext() bool
	// Get the (i, j) coordinate of the point visited by the iterator
	GetCoordinate() (int, int)
	// The iterator goes back to its initial position
	Reset()
}

type linearIterator struct {
	i    int
	wi   int
	he   int
	size int
}

// A Linear Iterator will simulate the behavior of a double for common course
func NewLinearIterator(mtrx *Matrix) (res Iterator) {
	res = &linearIterator{0, mtrx.width, mtrx.height, mtrx.width * mtrx.height}
	return
}

func (it *linearIterator) Next() {
	it.i++
}

func (it *linearIterator) HasNext() bool {
	return it.i != it.size
}

func (it *linearIterator) GetCoordinate() (i, j int) {
	i, j = it.i%it.wi, it.i/it.he
	return
}

func (it *linearIterator) Reset() {
	it.i = 0
}
