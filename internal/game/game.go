package game

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/eiannone/keyboard"
)

type Life struct {
	window Window
	pos    Position
	grid   *Grid
}

func NewGame(width, height, posX, posY int) *Life {
	const GRIDSIZE = 128

	window := NewWindow(int32(width), int32(height))
	pos := NewPosition(int32(posX), int32(posY))
	grid := NewGrid(GRIDSIZE)

	return &Life{
		window: window,
		pos:    pos,
		grid:   grid,
	}
}

func (g *Life) WithGrid(grid [][]byte, offsetX, offsetY int) *Life {
	for i, row := range grid {
		for j, cell := range row {
			g.grid.SetCell(i+offsetY, j+offsetX, int(cell))
		}
	}
	return g
}

func (g *Life) Run(tick time.Duration) {
	writer := bufio.NewWriter(os.Stdout)
	fmt.Println("\033[?25l\033[2J")
	defer fmt.Println("\033[?25h")
	keyChan, err := keyboard.GetKeys(10)
	if err != nil {
		log.Fatal("Unable to create key channel ", err)
	}
	defer keyboard.Close()
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case key := <-keyChan:
				if key.Key == keyboard.KeyEsc {
					cancel()
					return
				}
				switch key.Rune {
				case 'k':
					fallthrough
				case 'w':
					g.pos.Up()
				case 'h':
					fallthrough
				case 'a':
					g.pos.Left()
				case 'j':
					fallthrough
				case 's':
					g.pos.Down()
				case 'l':
					fallthrough
				case 'd':
					g.pos.Right()
				}
			default:
			}
		}
	}()
	go func() {
		ticker := time.NewTicker(tick)
		defer ticker.Stop()
		for range ticker.C {
			select {
			case <-ctx.Done():
				return
			default:
			}
			writer.Write([]byte("\033[H"))
			writer.Write(g.Serialize())
			writer.Flush()
			g.grid.Update()
		}
	}()
	<-ctx.Done()
}

func (g *Life) Serialize() []byte {
	posX := g.pos.GetPosX()
	posY := g.pos.GetPosY()
	width := g.window.GetWidth()
	height := g.window.GetHeight()
	out := make([]byte, 0, 10)
	for range width + 2 {
		out = append(out, '\033', '[', '3', '4', 'm')
		out = append(out, '\033', '[', '4', '4', 'm')
		out = append(out, 226, 172, 155, '\033', '[', '0', 'm')
	}
	out = append(out, '\n')
	for i := range height {
		out = append(out, '\033', '[', '3', '4', 'm')
		out = append(out, '\033', '[', '4', '4', 'm')
		out = append(out, 226, 172, 155, '\033', '[', '0', 'm')
		for j := range width {
			x, y := int(posX+j), int(posY+i)
			if y < 0 || y >= g.grid.Size() || x < 0 || x >= g.grid.Size() {
				out = append(out, '\033', '[', '3', '1', 'm')
				out = append(out, '\033', '[', '4', '1', 'm')
			} else if g.grid.GetCell(y, x) != 0 {
				out = append(out, '\033', '[', '3', '0', 'm')
				out = append(out, '\033', '[', '4', '0', 'm')
			} else {
				out = append(out, '\033', '[', '3', '7', 'm')
				out = append(out, '\033', '[', '4', '7', 'm')
			}
			out = append(out, 226, 172, 155, '\033', '[', '0', 'm')
		}
		out = append(out, '\033', '[', '3', '4', 'm')
		out = append(out, '\033', '[', '4', '4', 'm')
		out = append(out, 226, 172, 155, '\033', '[', '0', 'm')
		out = append(out, '\n')
	}
	for range width + 2 {
		out = append(out, '\033', '[', '3', '4', 'm')
		out = append(out, '\033', '[', '4', '4', 'm')
		out = append(out, 226, 172, 155, '\033', '[', '0', 'm')
	}
	return out
}
