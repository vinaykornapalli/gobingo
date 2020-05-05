package api

import (
	     "github.com/gorilla/mux" 
		 "net/http"
		)

//InitAPI to initialize server
func InitAPI() {
	r:=mux.NewRouter()
	r.HandleFunc("/newgame" , NewGame)
	r.HandleFunc("/joingame/{id:[0-9]+}" , JoinGame)
	http.ListenAndServe(":8080" ,r )
}


