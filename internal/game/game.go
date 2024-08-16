package game

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/eiannone/keyboard"

	"game_of_life/pkg/utils"
)

type Life struct {
	window Window
	grid   [][]byte
	pos    Position
	gridMu sync.RWMutex
}

func NewGame(width, height, posX, posY int) *Life {
	const GRIDSIZE = 128

	window := NewWindow(int32(width), int32(height))
	pos := NewPosition(int32(posX), int32(posY))

	grid := make([][]byte, GRIDSIZE)
	for i := range grid {
		grid[i] = make([]byte, GRIDSIZE)
	}

	return &Life{
		window: window,
		pos:    pos,
		grid:   grid,
		gridMu: sync.RWMutex{},
	}
}

func (g *Life) WithGrid(grid [][]byte, offsetX, offsetY int) *Life {
	for i, row := range grid {
		for j, cell := range row {
			g.grid[i+offsetY][j+offsetX] = cell
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
			g.grid = g.NewCycle()
		}
	}()
	<-ctx.Done()
}

func (g *Life) SetCell(row, col, state int) {
	utils.Assertf(row >= 0 && row < len(g.grid), "Invalid grid row: %d", row)
	utils.Assertf(col >= 0 && col < len(g.grid[row]), "Invalid grid column: %d", col)

	g.grid[row][col] = byte(state)
}

func (g *Life) GetCell(row, col int) byte {
	utils.Assertf(row >= 0 && row < len(g.grid), "Invalid grid row: %d", row)
	utils.Assertf(col >= 0 && col < len(g.grid[row]), "Invalid grid column: %d", col)

	return g.grid[row][col]
}

func (g *Life) NewCycle() [][]byte {
	const CHUNKSIZE = 56

	gridSize := len(g.grid)
	chunkCount := gridSize / CHUNKSIZE

	newGrid := make([][]byte, gridSize)

	rowWg := sync.WaitGroup{}
	rowWg.Add(gridSize)

	for rowIndex, row := range g.grid {
		go func() {
			defer rowWg.Done()
			chunks := make([][]byte, chunkCount)
			for chunkIndex := range chunkCount {
				chunks[chunkIndex] = row[chunkIndex*CHUNKSIZE : chunkIndex*CHUNKSIZE+CHUNKSIZE]
			}
			newRow := make([]byte, gridSize)
			chunkWg := sync.WaitGroup{}
			chunkWg.Add(chunkCount)
			for chunkIndex, chunk := range chunks {
				go func() {
					defer chunkWg.Done()
					for chunkOffset, cell := range chunk {
						colIndex := chunkIndex*CHUNKSIZE + chunkOffset
						alive, _ := g.CountNeighbor(rowIndex, colIndex)

						underPopulation := alive < 2
						overPopulation := alive > 3

						isAlive := cell != 0
						survival := isAlive && (alive == 2 || alive == 3)
						reproduce := !isAlive && alive == 3

						if underPopulation || overPopulation {
							newRow[colIndex] = 0
						} else if survival || reproduce {
							newRow[colIndex] = 1
						}

					}
				}()
			}
			chunkWg.Wait()
			newGrid[rowIndex] = newRow
		}()
	}
	rowWg.Wait()

	return newGrid
}

func (g *Life) CountNeighbor(row, col int) (alive int, dead int) {
	utils.Assertf(row >= 0 && row < len(g.grid), "Invalid grid row: %d", row)
	utils.Assertf(col >= 0 && col < len(g.grid[row]), "Invalid grid column: %d", col)

	alive, dead = 0, 0
	around := [][2]int{
		{col - 1, row - 1},
		{col, row - 1},
		{col + 1, row - 1},
		{col - 1, row},
		{col + 1, row},
		{col - 1, row + 1},
		{col, row + 1},
		{col + 1, row + 1},
	}
	g.gridMu.RLock()
	for _, coor := range around {
		x, y := coor[0], coor[1]
		if x < 0 || x >= len(g.grid[0]) || y < 0 || y >= len(g.grid) {
			continue
		}
		if g.grid[y][x] != 0 {
			alive++
			continue
		}
		dead++
	}
	g.gridMu.RUnlock()
	return alive, dead
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
			if y < 0 || y >= len(g.grid) || x < 0 || x >= len(g.grid[0]) {
				out = append(out, '\033', '[', '3', '1', 'm')
				out = append(out, '\033', '[', '4', '1', 'm')
			} else if g.grid[y][x] != 0 {
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
