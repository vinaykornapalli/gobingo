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
	p.bingoLines = 0

	return p
}

func (p *Player) updateMatrix(val int) {

	for i := 0; i < 5; i++ {

		for j := 0; j < 5; j++ {

			if p.playerMatrix.cell[i][j] == val {
				p.playerMatrix.cell[i][j] = 0
				p.updateBingoLines(i, j)
				break
			}
		}
	}

}

func (p *Player) updateBingoLines(row int, col int) {

	//checking row
	l := 0

	for i := 0; i < 5; i++ {

		if p.playerMatrix.cell[row][i] != 0 {
			l = 1
			break
		}
	}

	if l == 0 {
		for i := 0; i < 5; i++ {

			p.playerMatrix.cell[row][i] = 26
		}
	}

	p.bingoLines = p.bingoLines + 1
	p.playerMatrix.cell[row][col] = 0

	//checking column
	l = 0
	for i := 0; i < 5; i++ {

		if p.playerMatrix.cell[i][col] != 0 {
			l = 1
			break
		}
	}

	if l == 0 {
		for i := 0; i < 5; i++ {

			p.playerMatrix.cell[i][col] = 26
		}
	}

	p.bingoLines = p.bingoLines + 1
	p.playerMatrix.cell[row][col] = 0

	//checking diagonal1
	l = 0
	if row == col {

		for i := 0; i < 5; i++ {

			if p.playerMatrix.cell[i][i] != 0 {
				l = 1
				break
			}
		}
	}

	if l == 0 {
		for i := 0; i < 5; i++ {

			p.playerMatrix.cell[i][i] = 26
		}
	}

	p.bingoLines = p.bingoLines + 1
	p.playerMatrix.cell[row][col] = 0

	//checking diagonal2
	l = 0
	if row == 4-col {

		for i := 0; i < 5; i++ {

			if p.playerMatrix.cell[i][4-i] != 0 {
				l = 1
				break
			}
		}
	}

	if l == 0 {
		for i := 0; i < 5; i++ {

			p.playerMatrix.cell[i][4-i] = 26
		}
	}

	p.bingoLines = p.bingoLines + 1
}

func (p *Player) checkIsBingo() {
	if p.bingoLines == 5 {
		p.isBingo = true
	}
}
