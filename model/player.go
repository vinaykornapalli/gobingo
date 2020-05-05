package model

import (
	"github.com/google/uuid"
)

type Player struct {
	PlayerID     string
	Name         string
	PlayerMatrix BingoMatrix
	BingoLines   int
	IsBingo      bool
}

func CreatePlayer(pName string) Player {

	var p Player
	p.InitPlayer(pName)

	return p

}

func (p *Player) InitPlayer(pName string) {

	id := uuid.New()
	p.PlayerID = id.String()
	p.Name = pName
	p.PlayerMatrix.InitMatrix()
	p.BingoLines = 0
}

func (p *Player) updateBingoLines(chosenNumber int) {

	row, col := p.PlayerMatrix.UpdateMatrix(chosenNumber)

	//checking row
	l := 0

	for i := 0; i < 5; i++ {

		if p.PlayerMatrix.cell[row][i] != 0 {
			l = 1
			break
		}
	}

	if l == 0 {
		for i := 0; i < 5; i++ {

			p.PlayerMatrix.cell[row][i] = 26
		}
	}

	p.BingoLines = p.BingoLines + 1
	p.PlayerMatrix.cell[row][col] = 0

	//checking column
	l = 0
	for i := 0; i < 5; i++ {

		if p.PlayerMatrix.cell[i][col] != 0 {
			l = 1
			break
		}
	}

	if l == 0 {
		for i := 0; i < 5; i++ {

			p.PlayerMatrix.cell[i][col] = 26
		}
	}

	p.BingoLines = p.BingoLines + 1
	p.PlayerMatrix.cell[row][col] = 0

	//checking diagonal1
	l = 0
	if row == col {

		for i := 0; i < 5; i++ {

			if p.PlayerMatrix.cell[i][i] != 0 {
				l = 1
				break
			}
		}
	}

	if l == 0 {
		for i := 0; i < 5; i++ {

			p.PlayerMatrix.cell[i][i] = 26
		}
	}

	p.BingoLines = p.BingoLines + 1
	p.PlayerMatrix.cell[row][col] = 0

	//checking diagonal2
	l = 0
	if row == 4-col {

		for i := 0; i < 5; i++ {

			if p.PlayerMatrix.cell[i][4-i] != 0 {
				l = 1
				break
			}
		}
	}

	if l == 0 {
		for i := 0; i < 5; i++ {

			p.PlayerMatrix.cell[i][4-i] = 26
		}
	}

	p.BingoLines = p.BingoLines + 1
}

func (p *Player) checkIsBingo() {
	if p.BingoLines == 5 {
		p.IsBingo = true
	}
}
