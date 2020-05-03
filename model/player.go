package model

import (
	"github.com/google/uuid"
)

type Player struct {
	playerID     string
	name         string
	playerMatrix BingoMatrix
	bingoLines   int
	isBingo      bool
}

func initPlayer(pName string) Player {

	var p Player
	id := uuid.New()
	p.playerID = id.String()
	p.name = pName
	p.playerMatrix = initMatrix()

	return p
}

func (p *Player) updateMatrix(val int) {

}

func (p *Player) selectValue() {

}
