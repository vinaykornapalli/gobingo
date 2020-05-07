package api

import(
	"github.com/JamesStewy/sse"
	"net/http"
	"fmt"
	"time"
)


//GameSSE for server sent events for clients
type GameSSE struct {
	clients map[string]map[*sse.Client]bool
}

//creates client datastructure
func (gs *GameSSE) initSSE(){
  gs.clients = make(map[string]map[*sse.Client]bool)
} 

func (gs *GameSSE) eventHandler(w http.ResponseWriter, r *http.Request) {
	// Initialise (REQUIRED)
	ids , _ := r.URL.Query()["id"]
	id :=ids[0]
	client, err := sse.ClientInit(w)
	fmt.Println(id)
	// Return error if unable to initialise Server-Sent Events
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Add client to external variable for later use
	gs.clients[id] = make(map[*sse.Client]bool)
	gs.clients[id][client] = true
	// Remove client from external variable on exit
	// defer delete(clients, client)
    fmt.Println(client)
	// Run Client (REQUIRED)
	client.Run(r.Context())
}


//SendDataToClients sends events to clents of a specific game
func (gs *GameSSE)SendDataToClients(gameid string , data string) {
  
	clients := gs.clients[gameid]
  
	// msg:= sse.Msg{
	// 	Event:"game state",
	// 	Data: data,
	// }
	ticker := time.NewTicker(time.Second * 2)
	
	for t := range ticker.C {
		// Create new message called 'time' with data containing the current time
		fmt.Println(len(clients))
		msg := sse.Msg{
			Event: "time",
			Data:  t.String(),
		}

		for client := range clients {
			// Send the message to this client
			client.Send(msg)
		}
	}
	
}