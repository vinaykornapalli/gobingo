package api

import (
		 "github.com/gorilla/mux" 
		 "github.com/gorilla/handlers"
		 "net/http"
		)

var gs GameSSE 
//InitAPI to initialize server
func InitAPI() {
	
	
	gs.initSSE()
	go gs.SendDataToClients("1234" , "hey its workinggg")
	r:=mux.NewRouter()
	r.HandleFunc("/newgame" , NewGame)
	r.HandleFunc("/joingame/{id:[0-9]+}" , JoinGame)
	r.HandleFunc("/event", gs.eventHandler)
	http.ListenAndServe(":8080" ,handlers.CORS()(r) )
	
}


