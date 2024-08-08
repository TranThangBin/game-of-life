package test

import (
	"testing"

	internal "game_of_life/internal/game"
)

func TestMid(t *testing.T) {
	g := internal.NewGame(3, 3)
	g.SetCell(1, 1, 1)
	g.SetCell(2, 1, 1)
	// . . .
	// .|*|.
	// . * .
	alive, dead := g.CountNeighbor(1, 1)
	expectedAlive, expectedDead := 1, 7
	if alive != expectedAlive {
		t.Fatalf("Expected %d alive cell but get %d", expectedAlive, alive)
	}
	if dead != expectedDead {
		t.Fatalf("Expected %d dead cell but get %d", expectedDead, dead)
	}
}

func TestTopStart(t *testing.T) {
	g := internal.NewGame(3, 3)
	g.SetCell(0, 0, 1)
	g.SetCell(1, 1, 1)
	g.SetCell(2, 0, 1)
	g.SetCell(2, 2, 1)
	//|*|. .
	// . * .
	// * . *
	alive, dead := g.CountNeighbor(0, 0)
	expectedAlive, expectedDead := 1, 2
	if alive != expectedAlive {
		t.Fatalf("Expected %d alive cell but get %d", expectedAlive, alive)
	}
	if dead != expectedDead {
		t.Fatalf("Expected %d dead cell but get %d", expectedDead, dead)
	}
}
func TestTopEnd(t *testing.T) {
	g := internal.NewGame(3, 3)
	g.SetCell(0, 2, 1)
	g.SetCell(1, 0, 1)
	g.SetCell(1, 1, 1)
	g.SetCell(2, 0, 1)
	g.SetCell(2, 2, 1)
	// . .|*|
	// * * .
	// * . *
	alive, dead := g.CountNeighbor(0, 2)
	expectedAlive, expectedDead := 1, 2
	if alive != expectedAlive {
		t.Fatalf("Expected %d alive cell but get %d", expectedAlive, alive)
	}
	if dead != expectedDead {
		t.Fatalf("Expected %d dead cell but get %d", expectedDead, dead)
	}
}

func TestMidStart(t *testing.T) {
	g := internal.NewGame(3, 3)
	g.SetCell(0, 0, 1)
	g.SetCell(1, 0, 1)
	g.SetCell(2, 0, 1)
	g.SetCell(1, 1, 1)
	g.SetCell(1, 2, 1)
	// * . .
	//|*|* *
	// * . .
	alive, dead := g.CountNeighbor(1, 0)
	expectedAlive, expectedDead := 3, 2
	if alive != expectedAlive {
		t.Fatalf("Expected %d alive cell but get %d", expectedAlive, alive)
	}
	if dead != expectedDead {
		t.Fatalf("Expected %d dead cell but get %d", expectedDead, dead)
	}
}

func TestMidEnd(t *testing.T) {
	g := internal.NewGame(3, 3)
	g.SetCell(0, 2, 1)
	g.SetCell(1, 1, 1)
	g.SetCell(2, 0, 1)
	// . . *
	// . *|.|
	// * . .
	alive, dead := g.CountNeighbor(1, 2)
	expectedAlive, expectedDead := 2, 3
	if alive != expectedAlive {
		t.Fatalf("Expected %d alive cell but get %d", expectedAlive, alive)
	}
	if dead != expectedDead {
		t.Fatalf("Expected %d dead cell but get %d", expectedDead, dead)
	}
}

func TestBotStart(t *testing.T) {
	g := internal.NewGame(3, 3)
	g.SetCell(0, 2, 1)
	g.SetCell(1, 0, 1)
	g.SetCell(1, 1, 1)
	g.SetCell(2, 1, 1)
	// . . *
	// * * .
	//|.|* .
	alive, dead := g.CountNeighbor(2, 0)
	expectedAlive, expectedDead := 3, 0
	if alive != expectedAlive {
		t.Fatalf("Expected %d alive cell but get %d", expectedAlive, alive)
	}
	if dead != expectedDead {
		t.Fatalf("Expected %d dead cell but get %d", expectedDead, dead)
	}
}
func TestBotEnd(t *testing.T) {
	g := internal.NewGame(3, 3)
	g.SetCell(0, 2, 1)
	g.SetCell(1, 0, 1)
	g.SetCell(1, 1, 1)
	g.SetCell(2, 1, 1)
	// . . *
	// . * .
	// . *|*|
	alive, dead := g.CountNeighbor(2, 0)
	expectedAlive, expectedDead := 3, 0
	if alive != expectedAlive {
		t.Fatalf("Expected %d alive cell but get %d", expectedAlive, alive)
	}
	if dead != expectedDead {
		t.Fatalf("Expected %d dead cell but get %d", expectedDead, dead)
	}
}
