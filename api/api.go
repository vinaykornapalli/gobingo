package api

import (
		 "github.com/gorilla/mux" 
		 "github.com/gorilla/handlers"
		 "net/http"
		
		)

var data chan string
var streamID chan string
//InitAPI to initialize server
func InitAPI() {

	gs := NewSSE()
	data  = make(chan string)
	streamID = make(chan string)
	go gs.SendDataToClients(streamID , data)
	r:=mux.NewRouter()
	r.HandleFunc("/newgame" , NewGame)
	r.HandleFunc("/join/{id:[A-Za-z0-9-_]+}" , JoinGame)
	r.HandleFunc("/lobby/{id:[A-Za-z0-9-_]+}", Lobby)
	r.HandleFunc("/event", gs.eventHandler)
	http.ListenAndServe(":8080" ,handlers.CORS()(r) )
}

//UpdateSSEChannel is used to update SSE GameState
func UpdateSSEChannel(stream string , in string){
	 data <-in
	 streamID <-stream
}


