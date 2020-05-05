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
	GameID  string `json:"game_id"`
	State   GameState `json:"state"`
	Players []Player `json:"players"`
	Winner  Player	`json:"winner"`
}

//CreateNewGame creates and returns a new game
func CreateNewGame(playerName string) string {

	var g Game
	newGameID := g.InitGame(playerName)

	return newGameID
}

//InitGame intializes a new game
func (g *Game) InitGame(playerName string) string {
	id := uuid.New()
	g.GameID = id.String()
	g.AddPlayer(CreatePlayer(playerName))
	g.CreateGameStore()
	return g.GameID
}

//AddPlayer add a new player into an existing game
func (g *Game) AddPlayer(p Player) {
	g.Players = append(g.Players, p)
}

//StartGame startes the game once all players are ready
func (g *Game) StartGame() {
	g.State.CrntActivePlayer = 0
	g.State.MaxPlayers = len(g.Players)
}

//UpdateChosenNumber updates chosen number in state
func (g *Game) UpdateChosenNumber(val int) {
	g.State.ChosenNumber = val
}

//PerformGamechanges is a all in one game handler
func (g *Game) PerformGamechanges() {

	for i := range g.Players {
		fmt.Println(g.Players[i].PlayerMatrix)
		fmt.Println(g.State.ChosenNumber)
		g.Players[i].updateBingoLines(g.State.ChosenNumber)
		fmt.Println(g.Players[i].PlayerMatrix)
		if g.Players[i].IsBingo {

			g.Winner = g.Players[i]
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



//DB PART


//CreateGameStore is used to create a storage for game data
func (g *Game) CreateGameStore() (int,error) {
	
	fName:= g.GameID + ".json"
	f , err := os.Create(fName)
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
	defer f.Close()
  return co , err
}

//RetriveGameFromStore retirves type game data from game store
func (g *Game) RetriveGameFromStore(id string) {
	fName :=  id + ".json"
	data , _ := ioutil.ReadFile(fName)
	json.Unmarshal(data,g)

}

//UpdateGameStore is used to make changes in game data inside store
func (g *Game) UpdateGameStore(id string)(int , error){
	fName :=  id + ".json"
	f , _ := os.Open(fName)

	blob , err :=json.Marshal(g)
	if err!=nil {
		panic(err)
	}
    //saving to file
	co , err := f.Write(blob)
	if err!=nil {
		panic(err)
	}
	defer f.Close()
	return co , err
}

//DeleteGameStore is to be used to delete game data once game is completed
func (g *Game) DeleteGameStore(id string){
	fName :=  id + ".json"
	os.Remove(fName)
}

