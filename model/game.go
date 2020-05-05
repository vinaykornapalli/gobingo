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
	gameID  string `json:game_id`
	state   GameState `json:state`
	players []Player `json:players`
	winner  Player	`json:winner`
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
	g.CreateGameStore()
	return g.gameID
}

func (g *Game) AddPlayer(p Player) {
	g.players = append(g.players, p)
}

func (g *Game) StartGame() {
	g.state.crntActivePlayer = 0
	g.state.maxPlayers = len(g.players)
}

func (g *Game) UpdatechosenNumber(val int) {
	g.state.chosenNumber = val
}

func (g *Game) PerformGamechanges() {

	for i := range g.players {
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

//ExitGame is used to end the game
func (g *Game) ExitGame() {
	fmt.Println("winner is ", g.winner.name)
}



//DB PART


//CreateGameStore is used to create a storage for game data
func (g *Game) CreateGameStore() (int,error) {
	
	fname:= "store/" + g.gameID + ".json"
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

