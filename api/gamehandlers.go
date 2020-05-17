package api

import (
	"encoding/json"
	"net/http"
	"github.com/gobingo/model"
	"github.com/gobingo/utils"
	uuid "github.com/satori/go.uuid"
)

//NewGame starts a new game by the first player
func NewGame(w http.ResponseWriter, r *http.Request) {
	//get player name
	playerName := r.URL.Query().Get("name")
	game := model.Game{}

	//Generating Game ID
	gid, _ := uuid.FromString(game.InitGame(playerName))
	gameID := utils.Encode(&gid)

	//Generating Player ID
	pid, _ := uuid.FromString(game.Players[0].PlayerID)
	playerID := utils.Encode(&pid)

	js := struct {
		GameID   string `json:"game_id"`
		PlayerID string `json:"player_id"`
	}{
		GameID:   gameID,
		PlayerID: playerID,
	}

	out, err := json.Marshal(js)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/xml")
	w.Write(out)
}

//JoinGame starts a new game by the first player
func JoinGame(w http.ResponseWriter, r *http.Request) {
	gid := r.URL.Query().Get("id")
	playerName := r.URL.Query().Get("name")
	ID, _ := utils.Decode(gid)
	gameID := ID.String()

	game := model.Game{}
	//Opening the game
	game.RetriveGameFromStore(gameID)

	player := model.Player{}
	player.InitPlayer(playerName)
	game.AddPlayer(player)

	//Generating Player ID
	pid, _ := uuid.FromString(player.PlayerID)
	playerID := utils.Encode(&pid)

	js := struct {
		GameID   string `json:"game_id"`
		PlayerID string `json:"player_id"`
	}{
		GameID:   gameID,
		PlayerID: playerID,
	}

	out, err := json.Marshal(js)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/xml")
	w.Write(out)

}
