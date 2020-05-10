package api

import (
	"fmt"
	"net/http"
	"github.com/JamesStewy/sse"
)

//GameSSE for server sent events for clients
type GameSSE struct {
	streams map[string]*Stream
}

//Stream is a instance of a specific game
type Stream struct {
	clients map[*sse.Client]bool
}

//NewSSE creates a new GameSSE
func NewSSE() *GameSSE {
	gs := &GameSSE{
		streams: make(map[string]*Stream),
	}
	return gs
}

func (gs *GameSSE) eventHandler(w http.ResponseWriter, r *http.Request) {
	// Initialise (REQUIRED)
	ids, _ := r.URL.Query()["id"]
	id := ids[0]
	client, err := sse.ClientInit(w)
	fmt.Println(id)
	// Return error if unable to initialise Server-Sent Events
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if gs.streams[id] == nil {
		stream := &Stream{}
		stream.clients = make(map[*sse.Client]bool)
		stream.clients[client] = true
		gs.streams[id] = stream

	} else {
		gs.streams[id].clients[client] = true
	}

	// Add client to external variable for later use

	// Remove client from external variable on exit
	defer delete(gs.streams[id].clients, client)

	// Run Client (REQUIRED)
	client.Run(r.Context())
}

//SendDataToClients sends events to clents of a specific game
func (gs *GameSSE) SendDataToClients(id chan string, data chan string) {

	clients := make(map[*sse.Client]bool)

	select {
	case d := <-data:
		gameid := <-id
		if gs.streams[gameid] != nil {
			clients = gs.streams[gameid].clients
			fmt.Println(d)
			msg := sse.Msg{
				Event: "time",
				Data:  d,
			}
			for client := range clients {
				// Send the message to this client
				client.Send(msg)
			}
		}

	}
	// Create new message called 'time' with data containing the current time

}
