package models

import (
	"errors"
	"fmt"
)

type Board struct {
	grid      [][]byte
	curPlayer byte
}

func NewBoard() *Board {
	board := &Board{
		grid: [][]byte{
			{0, 0, 0}, {0, 0, 0}, {0, 0, 0},
		},
		curPlayer: 1,
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
	newBoard.curPlayer = board.curPlayer
	return newBoard
}

func (board *Board) At(x, y byte) byte {
	return board.grid[y][x]
}

func (board *Board) Move(x, y byte) error {
	if board.grid[y][x] != 0 {
		return errors.New(fmt.Sprintf("The space (%d,%d) is not empty", x, y))
	}

	board.grid[y][x] = board.curPlayer
	if board.curPlayer == 1 {
		board.curPlayer = 2
	} else {
		board.curPlayer = 1
	}
	return nil
}
