package main

type NN struct {
	board                [7][6]int //0 empty, 1 P1, 2 P2
	moves                []int     //row played
	lastMoveX, lastMoveY int
	lastPlayer           int
}
