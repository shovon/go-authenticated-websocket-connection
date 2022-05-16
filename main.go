package main

import (
	"net/http"
	"strings"
	"wsauth/ws"

	"wsauth/messages/servermessages"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
)

var upgrader = websocket.Upgrader{}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/path", func(w http.ResponseWriter, r *http.Request) {
		// Set up the conection
		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		defer c.Close()

		// Grab the client ID
		clientId := strings.TrimSpace(r.URL.Query().Get("client_id"))
		if len(clientId) <= 0 {
			log.Info().Msg("The client did not supply a client ID. Returning without client ID nor tree ID")
			msg := servermessages.CreateClientError(
				servermessages.ErrorPayload{
					Title:  "Client ID has not been set",
					Detail: "Please set the relevant client ID, via the `client_id` query parameter",
				},
			)
			c.WriteJSON(msg)
		}

		conn := ws.NewWrapper(c)
		defer conn.Stop()

	})
}
