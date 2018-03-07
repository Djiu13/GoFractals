// Implementation of View interface
// It uses termbox-go, see https://github.com/nsf/termbox-go

package view

import (
	"fmt"
	"fracture/algorithm"
	"fracture/data"
	"termbox-go"
	"time"
)

const hsize = 31
const wsize = 81
const lsize = 30

type datas struct {
	ox, oy, precision float64
	width, height     int
}

type TermboxScreen struct {
	currentFractal datas
	logs           []string
	stop           chan int
	nbGen          int
	coloration     Coloration
}

func NewTermboxScreen(width, height int, coloration Coloration) *TermboxScreen {
	return &TermboxScreen{datas{0.0, 0.0, 4.0 / float64(wsize), width, height}, make([]string, 0), make(chan int, 1), 0, coloration}
}

func (t *TermboxScreen) Init() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}

	t.Log("init interface: done")
}

func (t *TermboxScreen) Close() {
	termbox.Close()
}

func (t *TermboxScreen) Log(str string) {
	for j := 0; j < hsize; j++ {
		for i := 0; i < wsize; i++ {
			termbox.SetCell(wsize+1+i, j, ' ', termbox.ColorWhite, termbox.ColorDefault)
		}
	}

	str = fmt.Sprintf("* %s", str)

	t.logs = append([]string{str}, t.logs...)
	current := hsize - 1
	for _, log := range t.logs {
		current -= (len(log) / lsize)
		for i, c := range log {
			line := current + i/lsize
			if line >= 0 {
				termbox.SetCell(wsize+1+i%30, line, c, termbox.ColorWhite, termbox.ColorDefault)
			}
		}
		current--
	}
}

func (t *TermboxScreen) EventLoop(routine algorithm.Routine) {

	routine(t, wsize, hsize, t.currentFractal.ox, t.currentFractal.oy, t.currentFractal.precision)
loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			compute := false
			if ev.Key == termbox.KeyCtrlC {
				break loop
			} else if ev.Key == termbox.KeyArrowDown {
				t.currentFractal.oy += t.currentFractal.precision
				compute = true
			} else if ev.Key == termbox.KeyArrowUp {
				t.currentFractal.oy -= t.currentFractal.precision
				compute = true
			} else if ev.Key == termbox.KeyArrowRight {
				t.currentFractal.ox += t.currentFractal.precision
				compute = true
			} else if ev.Key == termbox.KeyArrowLeft {
				t.currentFractal.ox -= t.currentFractal.precision
				compute = true
			} else if ev.Key == termbox.KeyCtrlR {
				t.currentFractal.precision *= 0.9
				compute = true
			} else if ev.Key == termbox.KeyCtrlT {
				t.currentFractal.precision *= 1.1
				compute = true
			} else if ev.Key == termbox.KeyCtrlS {
				t.nbGen += 1
				is := NewImageSaver(t.currentFractal.width, t.currentFractal.height, fmt.Sprintf("test%d.png", t.nbGen))
				is.Init()
				precision := float64(wsize) * t.currentFractal.precision / float64(t.currentFractal.width)
				t.Log(fmt.Sprintf("precision %f", precision))

				routine(is, t.currentFractal.width, t.currentFractal.height, t.currentFractal.ox, t.currentFractal.oy, precision)
				t.Log("test")
			}

			if compute {
				routine(t, wsize, hsize, t.currentFractal.ox, t.currentFractal.oy, t.currentFractal.precision)
			}
		}
	}
	t.stop <- 0
}

func (t *TermboxScreen) Listen(channel chan data.Pair, matrix *data.Matrix) {
	point, ok := <-channel

	for ok {
		val := matrix.At(point.First, point.Second)

		c, _ := t.coloration(val).(rune)

		termbox.SetCell(point.First, point.Second, c, termbox.ColorWhite, termbox.ColorDefault)

		point, ok = <-channel
	}
}

func (t *TermboxScreen) FlushLoop() {

loop:
	for {
		select {
		case <-t.stop:
			break loop
		default:
			time.Sleep(1000000)
			t.drawInterface()
			termbox.Flush()
		}
	}
}

func (t *TermboxScreen) drawInterface() {
	for j := 0; j < hsize; j++ {
		termbox.SetCell(wsize, j, ' ', termbox.ColorWhite, termbox.ColorWhite)
	}
	for i := 0; i < wsize+lsize; i++ {
		termbox.SetCell(i, hsize, ' ', termbox.ColorWhite, termbox.ColorWhite)
	}
	test := fmt.Sprintf("x=%f   y=%f   precision=%f", t.currentFractal.ox, t.currentFractal.oy, t.currentFractal.precision)

	for i, c := range test {
		termbox.SetCell(i, hsize, c, termbox.ColorBlack, termbox.ColorWhite)
	}

	termbox.SetCell(wsize/2, hsize/2, ' ', termbox.ColorBlack, termbox.ColorWhite)
}
