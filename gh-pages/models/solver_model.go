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
	grid          string
	asPlayer, idx byte
}

func (solver *Solver) Score() map[byte]int8 {
	if solver.Board.GameOver {
		panic("Shouldn't happen!")
	}

	key := SolverCacheKey{grid: solver.Board.grid, asPlayer: solver.AsPlayer}
	scores := map[byte]int8{}
	for idx, cell := range solver.Board.grid {
		if cell == ' ' {
			x := idx % 3
			y := idx / 3
			child := Solver{
				Board:    solver.Board.Clone(),
				AsPlayer: solver.AsPlayer,
			}
			child.Board.Move(x, y)
			if child.Board.GameOver {
				var inc int8 = -100
				if child.Board.Winner == 0 {
					if child.Board.CurPlayer == solver.AsPlayer {
						inc = 1
					} else {
						inc = -1
					}
				} else if child.Board.Winner == solver.AsPlayer {
					inc = 100
				}
				scores[byte(idx)] = inc
			} else {
				key.idx = byte(idx)
				var childScore map[byte]int8
				if foundScore, found := solverCache[key]; !found {
					childScore = child.Score()
					solverCache[key] = childScore
				} else {
					childScore = foundScore
				}
				scores[byte(idx)] = it.Fold(maps.Values(childScore), op.Add, 0)
			}
		}
	}
	return scores
}
