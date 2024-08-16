package game

import (
	"sync/atomic"

	"game_of_life/pkg/utils"
)

type Window struct {
	width  int32
	height int32
}

func NewWindow(width, height int32) Window {
	utils.Assertf(width > 0, "Game width have to be greater than 0: %d", width)
	utils.Assertf(height > 0, "Game height have to be greater than 0: %d", height)

	return Window{
		width:  width,
		height: height,
	}
}

func (w *Window) GetWidth() int32 {
	return atomic.LoadInt32(&w.width)
}

func (w *Window) GetHeight() int32 {
	return atomic.LoadInt32(&w.height)
}
