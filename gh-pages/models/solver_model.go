package models

import (
	"github.com/BooleanCat/go-functional/v2/it"
	"github.com/BooleanCat/go-functional/v2/it/op"
	"maps"
	"math"
)

var solverCache = map[SolverCacheKey]map[byte]int{}

type Solver struct {
	Board           *Board
	AsPlayer, Level byte
}

type SolverCacheKey struct {
	grid                 string
	asPlayer, level, idx byte
}

func (solver *Solver) rawScore() map[byte]int {
	if solver.Board.GameOver {
		panic("Shouldn't happen!")
	}

	key := SolverCacheKey{grid: solver.Board.grid, asPlayer: solver.AsPlayer, level: solver.Level}
	scores := map[byte]int{}
	for idx, cell := range solver.Board.grid {
		if cell == ' ' {
			x := idx % 3
			y := idx / 3
			child := Solver{
				Board:    solver.Board.Clone(),
				AsPlayer: solver.AsPlayer,
				Level:    solver.Level - 1,
			}
			child.Board.Move(x, y)
			if child.Board.GameOver {
				var inc = -int(solver.Level)
				if child.Board.Winner == 0 {
					inc = 0
				} else if child.Board.Winner == solver.AsPlayer {
					inc = int(solver.Level)
				}
				scores[byte(idx)] = inc
			} else {
				key.idx = byte(idx)
				key.level = solver.Level - 1
				var childScore map[byte]int
				if foundScore, found := solverCache[key]; !found {
					childScore = child.rawScore()
					solverCache[key] = childScore
				} else {
					childScore = foundScore
				}
				sum := it.Fold(maps.Values(childScore), op.Add, 0)
				scores[byte(idx)] = sum
			}
		}
	}
	return scores
}

func (solver *Solver) Score() map[byte]float64 {
	rawScore := solver.rawScore()
	total := it.Fold(maps.Values(rawScore), func(agg float64, v int) float64 {
		return agg + math.Abs(float64(v))
	}, 0.0)

	return it.Fold2(maps.All(rawScore), func(agg map[byte]float64, k byte, v int) map[byte]float64 {
		agg[k] = float64(v) / total
		return agg
	}, map[byte]float64{})
}
