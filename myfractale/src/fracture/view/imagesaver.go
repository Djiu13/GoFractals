// Implementation of View interface
// It permits to save a fractal into a png file

package view

import (
	"fracture/algorithm"
	"fracture/data"
	"image"
	"image/color"
	"image/png"
	"os"
	"time"
)

type ImageSaver struct {
	img      *image.RGBA
	filename string
	file     *os.File
	stop chan int
}

func NewImageSaver(width, height int, file string) *ImageSaver {
	return &ImageSaver{image.NewRGBA(image.Rect(0, 0, width,
			height)), file, nil, make(chan int)}
}

func (is *ImageSaver) Init() {
	file, err := os.Create(is.filename)
	if err != nil {
		panic(err)
	}

	is.file = file
}

func (is *ImageSaver) Close() {
	defer func(){is.stop <- 1}()
	png.Encode(is.file, is.img)
	if err := is.file.Close(); err != nil {
		panic(err)
	}
}

func (is *ImageSaver) Log(str string) {
}

func (is *ImageSaver) EventLoop(routine algorithm.Routine) {
}

func (is *ImageSaver) Listen(channel chan data.Pair, matrix *data.Matrix) {
	point, ok := <-channel

	for ok {
		val := matrix.At(point.First, point.Second)

		/*var col Coloration = ASCII*/
		/*c, ok := col(val).(uint8)*/

		if val == -1 {
			val = 0
		}

		c := BlackAndWhite(val).(uint8)

		is.img.Set(point.First, point.Second, color.RGBA{c, c, c, 255})

		point, ok = <-channel
	}

	is.Close()
}

func (is *ImageSaver) FlushLoop() {
	for {
		if <-is.stop == 1{
			return
		} else {
			time.Sleep(100)
		}
	}
}
