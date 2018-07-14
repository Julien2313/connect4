package main

import (
	"fmt"
)

type Game struct {
	board                [7][6]int //0 empty, 1 P1, 2 P2
	moves                []int     //row played
	lastMoveX, lastMoveY int
	lastPlayer           int
}

func InitGame() *Game {
	g := &Game{
		moves:      make([]int, 0),
		lastPlayer: 2,
	}
	return g
}

func (g *Game) printBoard() {

	for l := 5; l >= 0; l-- {
		for r := 0; r < 7; r++ {
			if g.board[r][l] == 1 {
				fmt.Printf("0 ")
			} else if g.board[r][l] == 2 {
				fmt.Printf("X ")
			} else {
				fmt.Printf("_ ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (g *Game) countMoveForARow(row int) int {
	occurrence := 0

	for _, v := range g.moves {
		if v == row {
			occurrence++
		}
	}
	return occurrence
}

func (g *Game) PlayerToPlay() int {
	return (len(g.moves) % 2) + 1
}

func (g *Game) makeAMove(row int, line int) {
	g.lastPlayer = g.PlayerToPlay()

	g.lastMoveX, g.lastMoveY = row, line
	g.board[g.lastMoveX][g.lastMoveY] = g.lastPlayer
	g.moves = append(g.moves, row)
}

func (g *Game) CheckAndMakeAMove(row int) bool {
	if row < 0 || row > 7 {
		return false
	}
	line := 0
	if line = g.countMoveForARow(row); line == 6 {
		return false
	}

	g.makeAMove(row, line)

	return true
}

func (g *Game) GetLastMove() (int, int) {
	x := g.moves[len(g.moves)-1]

	return x, g.countMoveForARow(x)
}

func (g *Game) CheckIfWonLine() bool {
	connectInARow := 0

	for row := 0; row < NBRROWS; row++ {
		if g.board[row][g.lastMoveY] == g.lastPlayer {
			connectInARow++
			if connectInARow == 4 {
				return true
			}
		} else {
			connectInARow = 0
		}
	}
	return false
}

func (g *Game) CheckIfWonRow() bool {
	connectInALine := 0

	for line := 0; line < NBRLINES; line++ {
		if g.board[g.lastMoveX][line] == g.lastPlayer {
			connectInALine++
			if connectInALine == 4 {
				return true
			}
		} else {
			connectInALine = 0
		}
	}
	return false
}

func (g *Game) CheckIfWonDiagonal() bool {
	connectInADiagonal := 0
	boardTopX := g.countMoveForARow(g.lastMoveX) - 1

	x := max(0, g.lastMoveX-boardTopX)
	y := max(0, boardTopX-g.lastMoveX)

	for y < NBRLINES && x < NBRROWS {
		if g.board[x][y] == g.lastPlayer {
			connectInADiagonal++
			if connectInADiagonal == 4 {
				return true
			}
		} else {
			connectInADiagonal = 0
		}
		y++
		x++
	}

	connectInADiagonal = 0
	x = max(0, max(0, boardTopX+g.lastMoveX)-(NBRLINES-1))
	y = min(NBRLINES-1, max(0, boardTopX+g.lastMoveX))

	for y >= 0 && x < NBRROWS {
		if g.board[x][y] == g.lastPlayer {
			connectInADiagonal++
			if connectInADiagonal == 4 {
				return true
			}
		} else {
			connectInADiagonal = 0
		}
		y--
		x++
	}

	return false
}

func (g *Game) CheckIfWon() bool {

	if g.CheckIfWonLine() || g.CheckIfWonRow() || g.CheckIfWonDiagonal() {
		return true
	}
	return false
}
