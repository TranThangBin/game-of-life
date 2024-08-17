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
	err := keyboard.Open()
	if err != nil {
		log.Fatal("Unable to create key channel ", err)
	}
	defer keyboard.Close()
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			ch, key, err := keyboard.GetSingleKey()
			if err != nil {
				return
			}
			if key == keyboard.KeyEsc {
				cancel()
				return
			}
			switch ch {
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
	fmt.Println("\nSimulation over!")
}

func (g *Life) Serialize() []byte {
	posX := g.pos.GetPosX()
	posY := g.pos.GetPosY()
	width := g.window.GetWidth()
	height := g.window.GetHeight()
	out := make([]byte, 0)
	square := []byte{226, 172, 155}
	blueColorBuilder :=
		NewColorBuilderWithBytes(square).
			WithBgColor(BLUE).
			WithFgColor(BLUE)
	redColorBuilder :=
		NewColorBuilderWithBytes(square).
			WithBgColor(RED).
			WithFgColor(RED)
	whiteColorBuilder :=
		NewColorBuilderWithBytes(square).
			WithBgColor(WHITE).
			WithFgColor(WHITE)
	blackColorBuilder :=
		NewColorBuilderWithBytes(square).
			WithBgColor(BLACK).
			WithFgColor(BLACK)

	for range width + 2 {
		out = append(out, blueColorBuilder.Build()...)
	}
	out = append(out, '\n')
	for i := range height {
		out = append(out, blueColorBuilder.Build()...)
		for j := range width {
			x, y := int(posX+j), int(posY+i)
			if y < 0 || y >= g.grid.Size() || x < 0 || x >= g.grid.Size() {
				out = append(out, redColorBuilder.Build()...)
				continue
			}
			if g.grid.GetCell(y, x) != 0 {
				out = append(out, blackColorBuilder.Build()...)
				continue
			}
			out = append(out, whiteColorBuilder.Build()...)
			continue
		}
		out = append(out, blueColorBuilder.Build()...)
		out = append(out, '\n')
	}
	for range width + 2 {
		out = append(out, blueColorBuilder.Build()...)
	}
	return out
}
