package game

import "sync/atomic"

type Position struct {
	posX int32
	posY int32
}

func NewPosition(posX, posY int32) Position {
	return Position{
		posX: posX,
		posY: posY,
	}
}

func (p *Position) Up() {
	atomic.AddInt32(&p.posY, -1)
}

func (p *Position) Down() {
	atomic.AddInt32(&p.posY, 1)
}

func (p *Position) Left() {
	atomic.AddInt32(&p.posX, -1)
}

func (p *Position) Right() {
	atomic.AddInt32(&p.posX, 1)
}

func (p *Position) GetPosX() int32 {
	return atomic.LoadInt32(&p.posX)
}

func (p *Position) GetPosY() int32 {
	return atomic.LoadInt32(&p.posY)
}
