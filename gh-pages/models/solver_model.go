package models

import (
	"github.com/BooleanCat/go-functional/v2/it"
	"github.com/BooleanCat/go-functional/v2/it/op"
	"maps"
)

var solverCache = map[SolverCacheKey]map[byte]int8{}

type Solver struct {
	Board    *Board
	AsPlayer byte
}

type SolverCacheKey struct {
	g0_0, g1_0, g2_0 byte
	g0_1, g1_1, g2_1 byte
	g0_2, g1_2, g2_2 byte
	x, y             int
}

func BoardKey(board *Board, x, y int) SolverCacheKey {
	return SolverCacheKey{
		g0_0: board.grid[0][0],
		g1_0: board.grid[0][1],
		g2_0: board.grid[0][2],
		g0_1: board.grid[1][0],
		g1_1: board.grid[1][1],
		g2_1: board.grid[1][2],
		g0_2: board.grid[2][0],
		g1_2: board.grid[2][1],
		g2_2: board.grid[2][2],
		x:    x,
		y:    y,
	}
}

func (solver *Solver) Score() map[byte]int8 {
	if solver.Board.GameOver {
		panic("Shouldn't happen!")
	}

	key := BoardKey(solver.Board, 0, 0)
	scores := map[byte]int8{}
	for y, row := range solver.Board.grid {
		for x, cell := range row {
			idx := byte(x + y*3)
			if cell == 0 {
				child := Solver{
					Board:    solver.Board.Clone(),
					AsPlayer: solver.AsPlayer,
				}
				child.Board.Move(x, y)
				if child.Board.GameOver {
					var inc int8 = -2
					if child.Board.Winner == 0 {
						if child.Board.CurPlayer == solver.AsPlayer {
							inc = 1
						} else {
							inc = -1
						}
					} else if child.Board.Winner == solver.AsPlayer {
						inc = 2
					}
					scores[idx] = inc
				} else {
					key.x = x
					key.y = y
					var childScore map[byte]int8
					if foundScore, found := solverCache[key]; !found {
						childScore = child.Score()
						solverCache[key] = childScore
					} else {
						childScore = foundScore
					}
					scores[idx] = it.Fold(maps.Values(childScore), op.Add, 0)
				}
			}
		}
	}
	return scores
}
