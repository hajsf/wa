package waSocket

import (
	"wa/api"
)

const maxClients = 1

func init() {
	api.Passer = api.DataPasser{
		Data:       make(chan api.SSEData),
		Logs:       make(chan string),
		Connection: make(chan struct{}, maxClients),
	}
}
