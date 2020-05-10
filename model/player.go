package model

import (
	"fmt"

	"github.com/google/uuid"
)

//Player type
type Player struct {
	PlayerID     string       `json:"player_id"`
	Name         string       `json:"name"`
	PlayerMatrix BingoMatrix  `json:"player_matrix"`
	BingoLines   int          `json:"bingo_lines"`
	IsBingo      bool         `json:"is_bingo"`
}

//CreatePlayer creates and returns a new player
func CreatePlayer(pName string) Player {

	var p Player
	p.InitPlayer(pName)

	return p

}

//InitPlayer initilize a new player
func (p *Player) InitPlayer(pName string) {

	id := uuid.New()
	p.PlayerID = id.String()
	p.Name = pName
	p.PlayerMatrix.InitMatrix()
	p.BingoLines = 0
}

func (p *Player) updateBingoLines(ChosenNumber int) {

	row, col := p.PlayerMatrix.UpdateMatrix(ChosenNumber)
	fmt.Println(row, col)

	//checking row
	l := 0

	for i := 0; i < 5; i++ {

		if p.PlayerMatrix.Cell[row][i] != 0 {
			l = 1
			break
		}
	}

	if l == 0 {
		for i := 0; i < 5; i++ {

			p.PlayerMatrix.Cell[row][i] = 26
		}

		p.BingoLines = p.BingoLines + 1
		p.PlayerMatrix.Cell[row][col] = 0
	}

	//checking column
	l = 0
	for i := 0; i < 5; i++ {

		if p.PlayerMatrix.Cell[i][col] != 0 {
			l = 1
			break
		}
	}
	if l == 0 {
		for i := 0; i < 5; i++ {

			p.PlayerMatrix.Cell[i][col] = 26
		}

		p.BingoLines = p.BingoLines + 1
		p.PlayerMatrix.Cell[row][col] = 0
	}

	//checking diagonal1
	l = 0
	if row == col {

		for i := 0; i < 5; i++ {

			if p.PlayerMatrix.Cell[i][i] != 0 {
				l = 1
				break
			}
		}

		if l == 0 {
			for i := 0; i < 5; i++ {

				p.PlayerMatrix.Cell[i][i] = 26
			}

			p.BingoLines = p.BingoLines + 1
			p.PlayerMatrix.Cell[row][col] = 0
		}
	}

	//checking diagonal2
	l = 0

	if row == 4-col {

		for i := 0; i < 5; i++ {

			if p.PlayerMatrix.Cell[i][4-i] != 0 {
				l = 1
				break
			}
		}
		if l == 0 {
			for i := 0; i < 5; i++ {

				p.PlayerMatrix.Cell[i][4-i] = 26
			}
			p.BingoLines = p.BingoLines + 1
		}
	}

}

func (p *	Player) checkIsBingo() {
	if p.BingoLines == 5 {
		p.IsBingo = true
	}
}
