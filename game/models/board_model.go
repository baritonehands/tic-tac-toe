package models

import (
	"errors"
	"fmt"
	"github.com/BooleanCat/go-functional/v2/it"
	"slices"
)

type Board struct {
	grid                     string
	CurPlayer, Winner, Taken byte
	GameOver                 bool
}

func NewBoard() *Board {
	board := &Board{
		grid:      ".........",
		CurPlayer: 'X',
		Taken:     0,
	}
	return board
}

func CoordToIndex(x, y int) int {
	return x + y*3
}

func (board *Board) Clone() *Board {
	newBoard := NewBoard()
	newGrid := make([]byte, len(board.grid))
	for idx, cell := range board.grid {
		newGrid[idx] = byte(cell)
	}
	newBoard.CurPlayer = board.CurPlayer
	newBoard.grid = string(newGrid)
	return newBoard
}

func (board *Board) At(x, y int) byte {
	return board.grid[CoordToIndex(x, y)]
}

func (board *Board) trioWinner(indexes []byte) byte {
	var last *byte = nil
	for _, n := range indexes {
		cur := board.grid[n]

		if last != nil && *last != cur {
			return '.'
		}
		last = &cur
	}
	return *last
}

func (board *Board) Move(x, y int) error {
	idx := CoordToIndex(x, y)
	if board.grid[idx] != '.' {
		return errors.New(fmt.Sprintf("The space (%d,%d) is not empty", x, y))
	}

	if board.Winner != 0 {
		return errors.New(fmt.Sprintf("The game is over"))
	}

	newGrid := []byte(board.grid)
	newGrid[idx] = board.CurPlayer
	board.grid = string(newGrid)
	board.Taken = board.Taken + 1
	winner, found := board.computeWinner()
	board.Winner = winner
	board.GameOver = found

	if found {
		return nil
	}

	for _, cell := range board.grid {
		if cell == '.' {
			// Not a cat's game
			if board.CurPlayer == 'X' {
				board.CurPlayer = 'O'
			} else {
				board.CurPlayer = 'X'
			}
			return nil
		}
	}
	board.GameOver = true

	return nil
}

var combos = [][]byte{
	{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, // Rows
	{0, 3, 6}, {1, 4, 7}, {2, 5, 8}, // Cols
	{0, 4, 8}, {2, 4, 6}, // Diags
}

func (board *Board) computeWinner() (byte, bool) {
	comboWinners := it.Map(slices.Values(combos), board.trioWinner)
	return it.Find(comboWinners, func(winner byte) bool {
		return winner != '.'
	})
}
