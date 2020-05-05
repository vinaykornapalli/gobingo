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

func CreateNewGame(playerName string) Game {

	var g Game
	g.InitGame(playerName)

	return g
}

func (g *Game) InitGame(playerName string) {
	id := uuid.New()
	g.gameID = id.String()

	g.AddPlayer(CreatePlayer(playerName))
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

	for i, _ := range g.players {
		fmt.Println(g.players[i].playerMatrix)
		fmt.Println(g.state.chosenNumber)
		g.players[i].updateBingoLines(g.state.chosenNumber)
		fmt.Println(g.players[i].playerMatrix)
		if g.players[i].isBingo {

			g.winner = g.players[i]
			g.ExitGame()
			return
		}
	}

	g.state.UpdateState()

}

func (g *Game) ExitGame() {
	fmt.Println("Winner is ", g.winner.name)
}
