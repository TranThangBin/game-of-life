package game

import (
	"sync"

	"game_of_life/pkg/utils"
)

type Grid struct {
	grid [][]byte
	mu   sync.RWMutex
}

func NewGrid(size int) *Grid {
	grid := make([][]byte, size)
	for i := range grid {
		grid[i] = make([]byte, size)
	}
	return &Grid{
		grid: grid,
		mu:   sync.RWMutex{},
	}
}

func (g *Grid) SetCell(row, col, state int) {
	utils.Assertf(row >= 0 && row < len(g.grid), "Invalid grid row: %d", row)
	utils.Assertf(col >= 0 && col < len(g.grid[row]), "Invalid grid column: %d", col)

	g.grid[row][col] = byte(state)
}

func (g *Grid) GetCell(row, col int) byte {
	utils.Assertf(row >= 0 && row < len(g.grid), "Invalid grid row: %d", row)
	utils.Assertf(col >= 0 && col < len(g.grid[row]), "Invalid grid column: %d", col)

	return g.grid[row][col]
}

func (g *Grid) Size() int {
	return len(g.grid)
}

func (g *Grid) Update() {
	g.grid = g.NextGeneration()
}

func (g *Grid) NextGeneration() [][]byte {
	const CHUNKCOUNT = 8

	gridSize := len(g.grid)
	chunkSize := gridSize / CHUNKCOUNT

	newGrid := make([][]byte, gridSize)

	rowWg := sync.WaitGroup{}
	rowWg.Add(gridSize)

	for rowIndex, row := range g.grid {
		go func() {
			defer rowWg.Done()
			newRow := make([]byte, gridSize)
			chunkWg := sync.WaitGroup{}
			chunkWg.Add(CHUNKCOUNT)
			for chunkIndex := range CHUNKCOUNT {
				chunk := row[chunkIndex*chunkSize : chunkIndex*chunkSize+chunkSize]
				go func() {
					defer chunkWg.Done()
					for chunkOffset, cell := range chunk {
						colIndex := chunkIndex*chunkSize + chunkOffset
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

func (g *Grid) CountNeighbor(row, col int) (alive int, dead int) {
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
	g.mu.RLock()
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
	g.mu.RUnlock()
	return alive, dead
}
