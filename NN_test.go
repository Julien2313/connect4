package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitNN(t *testing.T) {
	g := InitGame()

	assert.Equal(t, 0, len(g.moves))

	for _, l := range g.board {
		for _, r := range l {
			assert.Equal(t, 0, r)
		}
	}
}
