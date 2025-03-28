package models

import (
	"fmt"
	"github.com/BooleanCat/go-functional/v2/it"
	"github.com/BooleanCat/go-functional/v2/it/op"
	"maps"
	"math"
	"strings"
)

var solverCache = map[SolverCacheKey]map[byte]int{}

type Solver struct {
	Board           *Board
	AsPlayer, Level byte
}

func NewSolver(board *Board) *Solver {
	return &Solver{
		Board:    board,
		AsPlayer: board.CurPlayer,
		Level:    9 - board.Taken,
	}
}

type SolverCacheKey struct {
	grid          string
	asPlayer, idx byte
}

func (solver *Solver) BoardUpdated() {
	solver.AsPlayer = solver.Board.CurPlayer
	solver.Level = 9 - solver.Board.Taken
}

func (solver *Solver) rawScore() map[byte]int {
	if solver.Board.GameOver {
		panic("Shouldn't happen!")
	}

	key := SolverCacheKey{grid: solver.Board.grid, asPlayer: solver.AsPlayer} //, level: solver.Level}
	scores := map[byte]int{}
	children := make(map[int]*Board, 9-solver.Board.Taken)
	anyWinner := false
	printScores := func() {
		padding := strings.Repeat("  ", int(solver.Level))
		fmt.Printf("%sP%c %d: %v\n%s%s\n%s%s\n%s%s\n", padding, solver.AsPlayer, solver.Level, scores,
			padding, solver.Board.grid[:3], padding, solver.Board.grid[3:6], padding, solver.Board.grid[6:])
	}
	for idx, cell := range solver.Board.grid {
		if cell == '.' {
			x := idx % 3
			y := idx / 3
			childBoard := solver.Board.Clone()
			childBoard.Move(x, y)
			children[idx] = childBoard

			fmt.Println(childBoard.grid)
			if childBoard.GameOver {

				if solver.Board.grid == "...OX...X" {
					fmt.Println(childBoard)
				}
				if childBoard.Winner == solver.AsPlayer {
					scores[byte(idx)] = 1
					anyWinner = true
				} else if childBoard.Winner != 0 {
					scores[byte(idx)] = -1
					anyWinner = true
				}
			}
		}
	}
	fmt.Printf("P%c: Short-circuit if: %v, %v\n%s\n%s\n%s\n", solver.AsPlayer, anyWinner, scores,
		solver.Board.grid[:3], solver.Board.grid[3:6], solver.Board.grid[6:])
	if anyWinner {
		printScores()
		// Player will always choose to win if possible
		return scores
	}

	clear(scores)
	for idx, childBoard := range children {
		childSolver := Solver{
			Board:    childBoard,
			AsPlayer: solver.AsPlayer,
			Level:    solver.Level - 1,
		}
		if childSolver.Board.GameOver {
			scores[byte(idx)] = 0
		} else {
			key.idx = byte(idx)
			//key.level = solver.Level - 1
			var childScore map[byte]int
			if foundScore, found := solverCache[key]; !found {
				childScore = childSolver.rawScore()
				solverCache[key] = childScore
			} else {
				childScore = foundScore
			}
			sum := it.Fold(maps.Values(childScore), op.Add, 0)
			scores[byte(idx)] = sum
		}
	}
	printScores()
	return scores
}

func (solver *Solver) Score() map[byte]float64 {
	rawScore := solver.rawScore()
	total := it.Fold(maps.Values(rawScore), func(agg float64, v int) float64 {
		return agg + math.Abs(float64(v))
	}, 0.0)

	if total == 0.0 {
		fmt.Println(rawScore)
		return map[byte]float64{}
	}

	return it.Fold2(maps.All(rawScore), func(agg map[byte]float64, k byte, v int) map[byte]float64 {
		agg[k] = float64(v) / total
		return agg
	}, map[byte]float64{})
}
