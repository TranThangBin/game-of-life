package game

// func TestCountNeighborMid(t *testing.T) {
// 	g := NewGame(3, 3)
// 	g.SetCell(1, 1, 1)
// 	g.SetCell(2, 1, 1)
// 	// . . .
// 	// .|*|.
// 	// . * .
// 	alive, dead := g.CountNeighbor(1, 1)
// 	expectedAlive, expectedDead := 1, 7
// 	if alive != expectedAlive {
// 		t.Fatalf("Expected %d alive cell but get %d", expectedAlive, alive)
// 	}
// 	if dead != expectedDead {
// 		t.Fatalf("Expected %d dead cell but get %d", expectedDead, dead)
// 	}
// }
//
// func TestCountNeighborTopStart(t *testing.T) {
// 	g := NewGame(3, 3)
// 	g.SetCell(0, 0, 1)
// 	g.SetCell(1, 1, 1)
// 	g.SetCell(2, 0, 1)
// 	g.SetCell(2, 2, 1)
// 	//|*|. .
// 	// . * .
// 	// * . *
// 	alive, dead := g.CountNeighbor(0, 0)
// 	expectedAlive, expectedDead := 1, 2
// 	if alive != expectedAlive {
// 		t.Fatalf("Expected %d alive cell but get %d", expectedAlive, alive)
// 	}
// 	if dead != expectedDead {
// 		t.Fatalf("Expected %d dead cell but get %d", expectedDead, dead)
// 	}
// }
// func TestCountNeighborTopEnd(t *testing.T) {
// 	g := NewGame(3, 3)
// 	g.SetCell(0, 2, 1)
// 	g.SetCell(1, 0, 1)
// 	g.SetCell(1, 1, 1)
// 	g.SetCell(2, 0, 1)
// 	g.SetCell(2, 2, 1)
// 	// . .|*|
// 	// * * .
// 	// * . *
// 	alive, dead := g.CountNeighbor(0, 2)
// 	expectedAlive, expectedDead := 1, 2
// 	if alive != expectedAlive {
// 		t.Fatalf("Expected %d alive cell but get %d", expectedAlive, alive)
// 	}
// 	if dead != expectedDead {
// 		t.Fatalf("Expected %d dead cell but get %d", expectedDead, dead)
// 	}
// }
//
// func TestCountNeighborMidStart(t *testing.T) {
// 	g := NewGame(3, 3)
// 	g.SetCell(0, 0, 1)
// 	g.SetCell(1, 0, 1)
// 	g.SetCell(2, 0, 1)
// 	g.SetCell(1, 1, 1)
// 	g.SetCell(1, 2, 1)
// 	// * . .
// 	//|*|* *
// 	// * . .
// 	alive, dead := g.CountNeighbor(1, 0)
// 	expectedAlive, expectedDead := 3, 2
// 	if alive != expectedAlive {
// 		t.Fatalf("Expected %d alive cell but get %d", expectedAlive, alive)
// 	}
// 	if dead != expectedDead {
// 		t.Fatalf("Expected %d dead cell but get %d", expectedDead, dead)
// 	}
// }
//
// func TestCountNeighborMidEnd(t *testing.T) {
// 	g := NewGame(3, 3)
// 	g.SetCell(0, 2, 1)
// 	g.SetCell(1, 1, 1)
// 	g.SetCell(2, 0, 1)
// 	// . . *
// 	// . *|.|
// 	// * . .
// 	alive, dead := g.CountNeighbor(1, 2)
// 	expectedAlive, expectedDead := 2, 3
// 	if alive != expectedAlive {
// 		t.Fatalf("Expected %d alive cell but get %d", expectedAlive, alive)
// 	}
// 	if dead != expectedDead {
// 		t.Fatalf("Expected %d dead cell but get %d", expectedDead, dead)
// 	}
// }
//
// func TestCountNeighborBotStart(t *testing.T) {
// 	g := NewGame(3, 3)
// 	g.SetCell(0, 2, 1)
// 	g.SetCell(1, 0, 1)
// 	g.SetCell(1, 1, 1)
// 	g.SetCell(2, 1, 1)
// 	// . . *
// 	// * * .
// 	//|.|* .
// 	alive, dead := g.CountNeighbor(2, 0)
// 	expectedAlive, expectedDead := 3, 0
// 	if alive != expectedAlive {
// 		t.Fatalf("Expected %d alive cell but get %d", expectedAlive, alive)
// 	}
// 	if dead != expectedDead {
// 		t.Fatalf("Expected %d dead cell but get %d", expectedDead, dead)
// 	}
// }
// func TestCountNeighborBotEnd(t *testing.T) {
// 	g := NewGame(3, 3)
// 	g.SetCell(0, 2, 1)
// 	g.SetCell(1, 0, 1)
// 	g.SetCell(1, 1, 1)
// 	g.SetCell(2, 1, 1)
// 	// . . *
// 	// . * .
// 	// . *|*|
// 	alive, dead := g.CountNeighbor(2, 0)
// 	expectedAlive, expectedDead := 3, 0
// 	if alive != expectedAlive {
// 		t.Fatalf("Expected %d alive cell but get %d", expectedAlive, alive)
// 	}
// 	if dead != expectedDead {
// 		t.Fatalf("Expected %d dead cell but get %d", expectedDead, dead)
// 	}
// }
//
// func TestNewCycleUnderPopulation(t *testing.T) {
// 	g := NewGame(3, 3)
// 	g.SetCell(0, 0, 1)
// 	g.SetCell(1, 1, 1)
// 	expectedGrid := [][]byte{
// 		{0, 0, 0},
// 		{0, 0, 0},
// 		{0, 0, 0},
// 	}
// 	gottenGrid := g.NewCycle()
//
// 	for i, row := range gottenGrid {
// 		for j, cell := range row {
// 			if cell != expectedGrid[i][j] {
// 				t.Fatalf("Mismatch cell expected %d, gotten %d: x = %d, y = %d",
// 					expectedGrid[i][j], cell, i, j)
// 			}
// 		}
// 	}
// }
//
// func TestNewCycleSurvive(t *testing.T) {
// 	g := NewGame(3, 3)
// 	g.SetCell(0, 0, 1)
// 	g.SetCell(1, 1, 1)
// 	g.SetCell(2, 2, 1)
// 	expectedGrid := [][]byte{
// 		{0, 0, 0},
// 		{0, 1, 0},
// 		{0, 0, 0},
// 	}
// 	gottenGrid := g.NewCycle()
//
// 	for y, row := range gottenGrid {
// 		for x, cell := range row {
// 			if cell != expectedGrid[y][x] {
// 				t.Fatalf("Unexpected cell expected %d, gotten %d: x = %d, y = %d",
// 					expectedGrid[y][x], cell, x, y)
// 			}
// 		}
// 	}
// }
//
// func TestNewCycleOverpopulation(t *testing.T) {
// 	g := NewGame(3, 3)
// 	g.SetCell(0, 0, 1)
// 	g.SetCell(0, 1, 1)
// 	g.SetCell(0, 2, 1)
// 	g.SetCell(1, 1, 1)
// 	g.SetCell(1, 2, 1)
// 	g.SetCell(2, 0, 1)
// 	g.SetCell(2, 2, 1)
// 	expectedGrid := [][]byte{
// 		{1, 0, 1},
// 		{0, 0, 0},
// 		{0, 0, 1},
// 	}
// 	gottenGrid := g.NewCycle()
//
// 	for y, row := range gottenGrid {
// 		for x, cell := range row {
// 			if cell != expectedGrid[y][x] {
// 				t.Fatalf("Unexpected cell expected %d, gotten %d: x = %d, y = %d",
// 					expectedGrid[y][x], cell, x, y)
// 			}
// 		}
// 	}
// }
//
// func TestNewCycleReproduce(t *testing.T) {
// 	g := NewGame(3, 3)
// 	g.SetCell(0, 0, 1)
// 	g.SetCell(1, 2, 1)
// 	g.SetCell(2, 1, 1)
// 	expectedGrid := [][]byte{
// 		{0, 0, 0},
// 		{0, 1, 0},
// 		{0, 0, 0},
// 	}
// 	gottenGrid := g.NewCycle()
//
// 	for y, row := range gottenGrid {
// 		for x, cell := range row {
// 			if cell != expectedGrid[y][x] {
// 				t.Fatalf("Unexpected cell expected %d, gotten %d: x = %d, y = %d",
// 					expectedGrid[y][x], cell, x, y)
// 			}
// 		}
// 	}
// }
