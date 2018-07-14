package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitGame(t *testing.T) {
	g := InitGame()

	assert.Equal(t, 0, len(g.moves))

	for _, l := range g.board {
		for _, r := range l {
			assert.Equal(t, 0, r)
		}
	}
}

func TestCheckAndMakeAMove(t *testing.T) {
	g := InitGame()

	assert.Equal(t, 0, len(g.moves))
	assert.Equal(t, true, g.CheckAndMakeAMove(0))
	assert.Equal(t, 1, g.board[0][0])

	assert.Equal(t, 1, len(g.moves))
	assert.Equal(t, true, g.CheckAndMakeAMove(0))
	assert.Equal(t, 2, g.board[0][1])

	assert.Equal(t, 2, len(g.moves))
	assert.Equal(t, true, g.CheckAndMakeAMove(0))
	assert.Equal(t, 1, g.board[0][2])

	assert.Equal(t, 3, len(g.moves))
	assert.Equal(t, true, g.CheckAndMakeAMove(0))
	assert.Equal(t, 2, g.board[0][3])

	assert.Equal(t, 4, len(g.moves))
	assert.Equal(t, true, g.CheckAndMakeAMove(0))
	assert.Equal(t, 1, g.board[0][4])

	assert.Equal(t, 5, len(g.moves))
	assert.Equal(t, true, g.CheckAndMakeAMove(0))
	assert.Equal(t, 2, g.board[0][5])

	assert.Equal(t, 6, len(g.moves))
	assert.Equal(t, false, g.CheckAndMakeAMove(0))
	assert.Equal(t, 2, g.board[0][5])

	assert.Equal(t, 6, len(g.moves))
}

func TestPlayerToPlay(t *testing.T) {
	g := InitGame()
	playerToPlay := 0

	for nbrMove := 0; nbrMove < 6; nbrMove++ {
		assert.Equal(t, playerToPlay+1, g.PlayerToPlay())
		assert.Equal(t, true, g.CheckAndMakeAMove(0))
		playerToPlay = 1 - playerToPlay
	}

}

func TestCheckIfWon(t *testing.T) {

	g := &Game{lastPlayer: 1,
		board: [7][6]int{
			[6]int{1, 0, 0, 0, 0, 0},
			[6]int{1, 0, 0, 0, 0, 0},
			[6]int{1, 0, 0, 0, 0, 0},
			[6]int{1, 0, 0, 0, 0, 0},
			[6]int{1, 0, 0, 0, 0, 0},
			[6]int{1, 0, 0, 0, 0, 0},
			[6]int{1, 0, 0, 0, 0, 0}},
		lastMoveX: 2,
		lastMoveY: 0,
		moves:     []int{0, 1, 2, 3, 4, 5, 6}}
	assert.Equal(t, true, g.CheckIfWonLine() && !g.CheckIfWonRow() && !g.CheckIfWonDiagonal())

	g = &Game{lastPlayer: 1,
		board: [7][6]int{
			[6]int{1, 0, 0, 0, 0, 0},
			[6]int{1, 0, 0, 0, 0, 0},
			[6]int{1, 0, 0, 0, 0, 0},
			[6]int{2, 0, 0, 0, 0, 0},
			[6]int{1, 0, 0, 0, 0, 0},
			[6]int{1, 0, 0, 0, 0, 0},
			[6]int{1, 0, 0, 0, 0, 0}},
		lastMoveX: 2,
		lastMoveY: 0,
		moves:     []int{0, 1, 2, 3, 4, 5, 6}}
	assert.Equal(t, true, !g.CheckIfWonLine() && !g.CheckIfWonRow() && !g.CheckIfWonDiagonal())

	g = &Game{lastPlayer: 1,
		board: [7][6]int{
			[6]int{1, 0, 0, 0, 0, 0},
			[6]int{1, 0, 0, 0, 0, 0},
			[6]int{1, 0, 0, 0, 0, 0},
			[6]int{1, 0, 0, 0, 0, 0},
			[6]int{0, 0, 0, 0, 0, 0},
			[6]int{0, 0, 0, 0, 0, 0},
			[6]int{0, 0, 0, 0, 0, 0}},
		lastMoveX: 2,
		lastMoveY: 0,
		moves:     []int{0, 1, 2, 3, 4}}
	assert.Equal(t, true, g.CheckIfWonLine() && !g.CheckIfWonRow() && !g.CheckIfWonDiagonal())

	g = &Game{lastPlayer: 2,
		board: [7][6]int{
			[6]int{0, 0, 0, 0, 0, 0},
			[6]int{0, 2, 0, 0, 0, 0},
			[6]int{0, 2, 0, 0, 0, 0},
			[6]int{0, 2, 0, 0, 0, 0},
			[6]int{0, 2, 0, 0, 0, 0},
			[6]int{0, 0, 0, 0, 0, 0},
			[6]int{0, 0, 0, 0, 0, 0}},
		lastMoveX: 2,
		lastMoveY: 1,
		moves:     []int{1, 2, 3, 4, 1, 2, 3, 4}}
	assert.Equal(t, true, g.CheckIfWonLine() && !g.CheckIfWonRow() && !g.CheckIfWonDiagonal())

	g = &Game{lastPlayer: 2,
		board: [7][6]int{
			[6]int{0, 0, 0, 0, 0, 0},
			[6]int{0, 2, 0, 0, 0, 0},
			[6]int{0, 2, 0, 0, 0, 0},
			[6]int{0, 1, 0, 0, 0, 0},
			[6]int{0, 2, 0, 0, 0, 0},
			[6]int{0, 2, 0, 0, 0, 0},
			[6]int{0, 0, 0, 0, 0, 0}},
		lastMoveX: 2,
		lastMoveY: 1,
		moves:     []int{1, 2, 3, 4, 1, 2, 3, 4, 5, 5}}
	assert.Equal(t, true, !g.CheckIfWonLine() && !g.CheckIfWonRow() && !g.CheckIfWonDiagonal())

	g = &Game{lastPlayer: 2,
		board: [7][6]int{
			[6]int{0, 0, 0, 0, 0, 0},
			[6]int{0, 2, 0, 0, 0, 0},
			[6]int{2, 2, 2, 2, 0, 0},
			[6]int{0, 1, 0, 0, 0, 0},
			[6]int{0, 2, 0, 0, 0, 0},
			[6]int{0, 2, 0, 0, 0, 0},
			[6]int{0, 0, 0, 0, 0, 0}},
		lastMoveX: 2,
		lastMoveY: 1,
		moves:     []int{1, 2, 3, 4, 1, 2, 3, 4, 5, 5, 2, 2}}
	assert.Equal(t, true, !g.CheckIfWonLine() && g.CheckIfWonRow() && !g.CheckIfWonDiagonal())

	g = &Game{lastPlayer: 2,
		board: [7][6]int{
			[6]int{0, 0, 0, 0, 0, 0},
			[6]int{0, 2, 0, 0, 0, 0},
			[6]int{2, 1, 2, 2, 0, 0},
			[6]int{0, 1, 0, 0, 0, 0},
			[6]int{0, 2, 0, 0, 0, 0},
			[6]int{0, 2, 0, 0, 0, 0},
			[6]int{0, 0, 0, 0, 0, 0}},
		lastMoveX: 2,
		lastMoveY: 1,
		moves:     []int{1, 2, 3, 4, 1, 2, 3, 4, 5, 5, 2, 2}}
	assert.Equal(t, true, !g.CheckIfWonLine() && !g.CheckIfWonRow() && !g.CheckIfWonDiagonal())

	g = &Game{lastPlayer: 2,
		board: [7][6]int{
			[6]int{0, 0, 0, 0, 0, 0},
			[6]int{0, 2, 0, 0, 0, 0},
			[6]int{1, 2, 2, 2, 2, 1},
			[6]int{0, 1, 0, 0, 0, 0},
			[6]int{0, 2, 0, 0, 0, 0},
			[6]int{0, 2, 0, 0, 0, 0},
			[6]int{0, 0, 0, 0, 0, 0}},
		lastMoveX: 2,
		lastMoveY: 1,
		moves:     []int{1, 2, 3, 4, 1, 2, 3, 4, 5, 5, 2, 2, 2, 2}}
	assert.Equal(t, true, !g.CheckIfWonLine() && g.CheckIfWonRow() && !g.CheckIfWonDiagonal())

	g = &Game{lastPlayer: 2,
		board: [7][6]int{
			[6]int{0, 0, 0, 0, 0, 0},
			[6]int{0, 2, 0, 0, 0, 0},
			[6]int{0, 0, 2, 0, 0, 0},
			[6]int{0, 0, 0, 2, 0, 0},
			[6]int{0, 0, 0, 0, 2, 0},
			[6]int{0, 0, 0, 0, 0, 0},
			[6]int{0, 0, 0, 0, 0, 0}},
		lastMoveX: 2,
		lastMoveY: 2,
		moves:     []int{1, 1, 2, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4, 4}}
	assert.Equal(t, true, !g.CheckIfWonLine() && !g.CheckIfWonRow() && g.CheckIfWonDiagonal())

	g = &Game{lastPlayer: 2,
		board: [7][6]int{
			[6]int{0, 0, 0, 2, 0, 0},
			[6]int{0, 0, 2, 0, 0, 0},
			[6]int{0, 2, 0, 0, 0, 0},
			[6]int{2, 0, 0, 0, 0, 0},
			[6]int{0, 0, 0, 0, 0, 0},
			[6]int{0, 0, 0, 0, 0, 0},
			[6]int{0, 0, 0, 0, 0, 0}},
		lastMoveX: 2,
		lastMoveY: 1,
		moves:     []int{0, 0, 0, 0, 1, 1, 1, 2, 2, 3}}
	assert.Equal(t, true, !g.CheckIfWonLine() && !g.CheckIfWonRow() && g.CheckIfWonDiagonal())
}
