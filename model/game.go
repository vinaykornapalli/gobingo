package model

import (
	"fmt"

	"github.com/google/uuid"
)

type Game struct {
	gameID  string
	state   GameState
	players []Player
	winner  Player
}

func CreateNewGame(playerName string) string {

	var g Game
	newgameID := g.InitGame(playerName)

	return newgameID
}

func (g *Game) InitGame(playerName string) string {
	id := uuid.New()
	g.gameID = id.String()

	g.AddPlayer(CreatePlayer(playerName))

	return g.gameID
}

func (g *Game) AddPlayer(p Player) {
	g.players = append(g.players, p)
}

func (g *Game) StartGame() {
	g.state.crntActivePlayer = 0
	g.state.maxPlayers = len(g.players)
}

func (g *Game) UpdateChosenNumber(val int) {
	g.state.chosenNumber = val
}

func (g *Game) PerformGamechanges() {

	for _, val := range g.players {

		val.updateBingoLines(g.state.chosenNumber)

		if val.isBingo {

			g.winner = val
			g.ExitGame()
			return
		}
	}

	g.state.UpdateState()

}

func (g *Game) ExitGame() {
	fmt.Println("Winner is ", g.winner.name)
}
