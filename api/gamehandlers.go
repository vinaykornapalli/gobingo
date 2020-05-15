package api

import (
	"fmt"
	"github.com/gorilla/mux" 
	"net/http"
	"github.com/gobingo/model"
	"encoding/json"
   )

//NewGame starts a new game by the first player
func NewGame(w http.ResponseWriter , r *http.Request) {
  name := r.URL.Query().Get("name")
  game := model.Game{}
  gameID := game.InitGame(name)
  
  js := struct{GameID string `json:"game_id"`}{GameID: gameID}
  res , err := json.Marshal(js)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  w.Header().Set("Content-Type", "application/json")
  w.Write(res)
}



//JoinGame starts a new game by the first player
func JoinGame(w http.ResponseWriter , r *http.Request) {
	vars := mux.Vars(r)
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Vars: %v\n", vars)
}