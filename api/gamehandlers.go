package api

import (
	"fmt"
	"github.com/gorilla/mux" 
	"net/http"
   )

//NewGame starts a new game by the first player
func NewGame(w http.ResponseWriter , r *http.Request) {

}

//JoinGame starts a new game by the first player
func JoinGame(w http.ResponseWriter , r *http.Request) {
	vars := mux.Vars(r)
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Vars: %v\n", vars)
}