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

func initGame(playerName string) string {

	var g Game

	id := uuid.New()
	g.gameID = id.String()
	g.addPlayer(initPlayer(playerName))

	return g.gameID

}

func (g *Game) addPlayer(p Player) {
	g.players = append(g.players, p)
}

func (g *Game) startGame() {
	g.state.crntActivePlayer = 0
	g.players[0].selectValue()
}

func (g *Game) performGamechanges() {

	for _, val := range g.players {

		val.updateMatrix(g.state.chosenNumber)

		if val.isBingo {

			g.winner = val
			g.exitGame()
			return
		}
	}

}

func (g *Game) exitGame() {
	fmt.Println("Winner is ", g.winner.name)
}
