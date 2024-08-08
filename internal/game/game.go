package internal

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"time"

	"game_of_life/pkg/utils"
)

type Life struct {
	grid [][]byte
}

func NewGame(width, height int) *Life {
	utils.Assertf(width > 0, "Game width have to be greater than 0: %d", width)
	utils.Assertf(height > 0, "Game height have to be greater than 0: %d", height)

	grid := make([][]byte, height)
	for i := range grid {
		grid[i] = make([]byte, width)
	}
	return &Life{grid: grid}
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
	ticker := time.NewTicker(tick)
	fmt.Println("\033[?25l\033[2J")
	defer fmt.Println("\033[?25h")
	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt)
	go func() {
		for {
			writer.Write([]byte("\033[H"))
			writer.Write(g.Serialize())
			writer.Flush()
			g.grid = g.NewCycle()
			<-ticker.C
		}
	}()
	<-s
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
	width := len(g.grid[0])
	height := len(g.grid)
	newGrid := make([][]byte, height)
	for i := range g.grid {
		newGrid[i] = make([]byte, width)
		for j, cell := range g.grid[i] {
			alive, _ := g.CountNeighbor(i, j)

			underPopulation := alive < 2
			overPopulation := alive > 3

			isAlive := cell != 0
			survival := isAlive && (alive == 2 || alive == 3)
			reproduce := !isAlive && alive == 3

			if underPopulation || overPopulation {
				newGrid[i][j] = 0
			} else if survival || reproduce {
				newGrid[i][j] = 1
			}
		}
	}
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
	return alive, dead
}

func (g *Life) Serialize() []byte {
	out := make([]byte, 0, 10)
	for _, row := range g.grid {
		for _, cell := range row {
			if cell != 0 {
				out = append(out, '\033', '[', '3', '0', 'm')
				out = append(out, '\033', '[', '4', '0', 'm')
			} else {
				out = append(out, '\033', '[', '3', '7', 'm')
				out = append(out, '\033', '[', '4', '7', 'm')
			}
			out = append(out, 226, 172, 155, '\033', '[', '0', 'm')
		}
		out = append(out, '\n')
	}
	return out
}
