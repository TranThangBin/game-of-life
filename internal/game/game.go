package game

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/eiannone/keyboard"

	"game_of_life/pkg/utils"
)

type Life struct {
	window [2]int
	pos    [2]int
	grid   [][]byte
	mu     sync.Mutex
}

func NewGame(width, height, posX, posY int) *Life {
	utils.Assertf(width > 0, "Game width have to be greater than 0: %d", width)
	utils.Assertf(height > 0, "Game height have to be greater than 0: %d", height)

	grid := make([][]byte, 512)
	for i := range grid {
		grid[i] = make([]byte, 512)
	}

	return &Life{
		window: [2]int{width, height},
		pos:    [2]int{posX, posY},
		grid:   grid,
		mu:     sync.Mutex{},
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
	ticker := time.NewTicker(tick)
	fmt.Println("\033[?25l\033[2J")
	defer fmt.Println("\033[?25h")
	keyChan, err := keyboard.GetKeys(10)
	if err != nil {
		log.Fatal("Unable to create key channel ", err)
	}
	defer keyboard.Close()
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case key := <-keyChan:
				if key.Key == keyboard.KeyEsc {
					return
				}
				switch key.Rune {
				case 'w':
					g.pos[1]--
				case 'a':
					g.pos[0]--
				case 's':
					g.pos[1]++
				case 'd':
					g.pos[0]++
				}
			default:
			}
			writer.Write([]byte("\033[H"))
			writer.Write(g.Serialize())
			writer.Flush()
			g.grid = g.NewCycle()
			<-ticker.C
		}
	}()
	wg.Wait()
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

	mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(height)

	for i := range g.grid {
		go func() {
			defer wg.Done()
			row := make([]byte, width)

			for j, cell := range g.grid[i] {
				alive, _ := g.CountNeighbor(i, j)

				underPopulation := alive < 2
				overPopulation := alive > 3

				isAlive := cell != 0
				survival := isAlive && (alive == 2 || alive == 3)
				reproduce := !isAlive && alive == 3

				if underPopulation || overPopulation {
					row[j] = 0
				} else if survival || reproduce {
					row[j] = 1
				}
			}

			mu.Lock()
			newGrid[i] = row
			mu.Unlock()
		}()
	}
	wg.Wait()

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
	g.mu.Lock()
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
	g.mu.Unlock()
	return alive, dead
}

func (g *Life) Serialize() []byte {
	out := make([]byte, 0, 10)
	for range g.window[0] + 2 {
		out = append(out, '\033', '[', '3', '4', 'm')
		out = append(out, '\033', '[', '4', '4', 'm')
		out = append(out, 226, 172, 155, '\033', '[', '0', 'm')
	}
	out = append(out, '\n')
	for i := range g.window[1] {
		out = append(out, '\033', '[', '3', '4', 'm')
		out = append(out, '\033', '[', '4', '4', 'm')
		out = append(out, 226, 172, 155, '\033', '[', '0', 'm')
		for j := range g.window[0] {
			x, y := g.pos[0]+j, g.pos[1]+i
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
	for range g.window[0] + 2 {
		out = append(out, '\033', '[', '3', '4', 'm')
		out = append(out, '\033', '[', '4', '4', 'm')
		out = append(out, 226, 172, 155, '\033', '[', '0', 'm')
	}
	return out
}
