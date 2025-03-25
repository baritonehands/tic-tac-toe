package models

import (
	"errors"
	"fmt"
	"github.com/BooleanCat/go-functional/v2/it"
	"slices"
)

type Board struct {
	grid              [][]byte
	CurPlayer, Winner byte
	GameOver          bool
}

func NewBoard() *Board {
	board := &Board{
		grid: [][]byte{
			{0, 0, 0}, {0, 0, 0}, {0, 0, 0},
		},
		CurPlayer: 1,
	}
	return board
}

func (board *Board) Clone() *Board {
	newBoard := NewBoard()
	for y, row := range board.grid {
		for x := range row {
			newBoard.grid[y][x] = board.grid[y][x]
		}
	}
	newBoard.CurPlayer = board.CurPlayer
	return newBoard
}

func (board *Board) At(x, y int) byte {
	return board.grid[y][x]
}

func (board *Board) trioWinner(idx []byte) byte {
	var last *byte = nil
	for _, n := range idx {
		x := int(n % 3)
		y := int(n / 3)
		cur := board.At(x, y)

		if last != nil && *last != cur {
			return 0
		}
		last = &cur
	}
	return *last
}

func (board *Board) CurPlayerName() string {
	if board.CurPlayer == 1 {
		return "X"
	} else {
		return "O"
	}
}

func (board *Board) Move(x, y int) error {
	if board.grid[y][x] != 0 {
		return errors.New(fmt.Sprintf("The space (%d,%d) is not empty", x, y))
	}

	if board.Winner != 0 {
		return errors.New(fmt.Sprintf("The game is over"))
	}

	board.grid[y][x] = board.CurPlayer
	winner, found := board.computeWinner()
	board.Winner = winner
	board.GameOver = found

	if found {
		return nil
	}

	for _, row := range board.grid {
		for _, col := range row {
			if col == 0 {
				// Not a cat's game
				if board.CurPlayer == 1 {
					board.CurPlayer = 2
				} else {
					board.CurPlayer = 1
				}
				return nil
			}
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
		return winner != 0
	})
}
