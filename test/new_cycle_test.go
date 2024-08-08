package test

import (
	"testing"

	internal "game_of_life/internal/game"
)

func TestUnderPopulation(t *testing.T) {
	g := internal.NewGame(3, 3)
	g.SetCell(0, 0, 1)
	g.SetCell(1, 1, 1)
	expectedGrid := [][]byte{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
	gottenGrid := g.NewCycle()

	for i, row := range gottenGrid {
		for j, cell := range row {
			if cell != expectedGrid[i][j] {
				t.Fatalf("Mismatch cell expected %d, gotten %d: x = %d, y = %d",
					expectedGrid[i][j], cell, i, j)
			}
		}
	}
}

func TestSurvive(t *testing.T) {
	g := internal.NewGame(3, 3)
	g.SetCell(0, 0, 1)
	g.SetCell(1, 1, 1)
	g.SetCell(2, 2, 1)
	expectedGrid := [][]byte{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	gottenGrid := g.NewCycle()

	for y, row := range gottenGrid {
		for x, cell := range row {
			if cell != expectedGrid[y][x] {
				t.Fatalf("Unexpected cell expected %d, gotten %d: x = %d, y = %d",
					expectedGrid[y][x], cell, x, y)
			}
		}
	}
}

func TestOverpopulation(t *testing.T) {
	g := internal.NewGame(3, 3)
	g.SetCell(0, 0, 1)
	g.SetCell(0, 1, 1)
	g.SetCell(0, 2, 1)
	g.SetCell(1, 1, 1)
	g.SetCell(1, 2, 1)
	g.SetCell(2, 0, 1)
	g.SetCell(2, 2, 1)
	expectedGrid := [][]byte{
		{1, 0, 1},
		{0, 0, 0},
		{0, 0, 1},
	}
	gottenGrid := g.NewCycle()

	for y, row := range gottenGrid {
		for x, cell := range row {
			if cell != expectedGrid[y][x] {
				t.Fatalf("Unexpected cell expected %d, gotten %d: x = %d, y = %d",
					expectedGrid[y][x], cell, x, y)
			}
		}
	}
}

func TestReproduce(t *testing.T) {
	g := internal.NewGame(3, 3)
	g.SetCell(0, 0, 1)
	g.SetCell(1, 2, 1)
	g.SetCell(2, 1, 1)
	expectedGrid := [][]byte{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	gottenGrid := g.NewCycle()

	for y, row := range gottenGrid {
		for x, cell := range row {
			if cell != expectedGrid[y][x] {
				t.Fatalf("Unexpected cell expected %d, gotten %d: x = %d, y = %d",
					expectedGrid[y][x], cell, x, y)
			}
		}
	}
}
