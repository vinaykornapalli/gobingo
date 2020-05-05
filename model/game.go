package model

import (
	"fmt"
	"encoding/json"
	"github.com/google/uuid"
	"os"
	"io/ioutil"
)
//Game Struct
type Game struct {
	GameID  string `json:game_id`
	State   GameState `json:state`
	Players []Player `json:players`
	Winner  Player	`json:winner`
}

func CreateNewGame(playerName string) string {

	var g Game
	newgameID := g.InitGame(playerName)

	return newgameID
}

func (g *Game) InitGame(playerName string) string {
	id := uuid.New()
	g.GameID = id.String()
	g.AddPlayer(CreatePlayer(playerName))
	g.CreateGameStore()
	return g.GameID
}

func (g *Game) AddPlayer(p Player) {
	g.Players = append(g.Players, p)
}

func (g *Game) StartGame() {
	g.State.CrntActivePlayer = 0
	g.State.MaxPlayers = len(g.Players)
}

func (g *Game) UpdateChosenNumber(val int) {
	g.State.ChosenNumber = val
}

func (g *Game) PerformGamechanges() {

	for _, val := range g.Players {

		val.updateBingoLines(g.State.ChosenNumber)

		if val.IsBingo {
			g.Winner = val
			g.ExitGame()
			return
		}
	}

	g.State.UpdateState()

}

//ExitGame is used to end the game
func (g *Game) ExitGame() {
	fmt.Println("Winner is ", g.Winner.Name)
}

//CreateGameStore is used to create a storage for game data
func (g *Game) CreateGameStore() (int,error) {
	
	fname:= "store/" + g.GameID + ".json"
	f , err := os.Create(fname)
	if err!=nil {
		panic(err)
	}

	blob , err :=json.Marshal(g)
	if err!=nil {
		panic(err)
	}
    //saving to file
	co , err := f.Write(blob)
	if err!=nil {
		panic(err)
	}

  return co , err
}

//RetriveGameFromStore retirves type game data from game store
func (g *Game) RetriveGameFromStore(id string) {
	fname := "store/" + id + ".json"
	data , _ := ioutil.ReadFile(fname)
	json.Unmarshal(data,g)

}

//UpdateGameStore is used to make changes in game data inside store
func (g *Game) UpdateGameStore(id string)(int , error){
	fname := "store/" + id + ".json"
	f , _ := os.Open(fname)

	blob , err :=json.Marshal(g)
	if err!=nil {
		panic(err)
	}
    //saving to file
	co , err := f.Write(blob)
	if err!=nil {
		panic(err)
	}
	return co , err
}

//DeleteGameStore is to be used to delete game data once game is completed
func (g *Game) DeleteGameStore(id string){
	fname := "store/" + id + ".json"
	os.Remove(fname)
}

