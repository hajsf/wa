package api

import (
	"fmt"
	"net/http"
)

type SSEData struct {
	Event, Message string
}
type DataPasser struct {
	Data       chan SSEData
	Logs       chan string
	Connection chan struct{} // To control maximum allowed clients connections
}

var Passer DataPasser

func (p DataPasser) HandleSignal(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream; charset=utf-8")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Cache-Control", "no-cache")
	// Allow cross origin  if required
	setupCORS(&w, r)

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Internal error", 500)
		return
	}

	fmt.Println("Client connected from IP:", r.RemoteAddr)
	// fmt.Println(len(p.connection), "new connection recieved")
	if len(p.Connection) > 0 {
		fmt.Fprint(w, "event: notification\ndata: Connection is opened in another browser/tap ...\n\n")
		flusher.Flush()
	}
	p.Connection <- struct{}{}

	fmt.Fprint(w, "event: notification\ndata: Connecting to WhatsApp server ...\n\n")
	flusher.Flush()

	// Connect to the WhatsApp client
	go p.Connect()

	for {
		select {
		case data := <-p.Data:
			// fmt.Println("SSE data recieved")
			//fmt.Println("msg recieved:", data.Message)
			switch {
			case len(data.Event) > 0:
				fmt.Fprintf(w, "event: %v\ndata: %v\n\n", data.Event, data.Message)
			case len(data.Event) == 0:
				fmt.Fprintf(w, "data: %v\n\n", data.Message)
			}
			flusher.Flush()
		case <-r.Context().Done():
			<-p.Connection
			fmt.Println("Connection closed from IP:", r.RemoteAddr)
			return
		}
	}
}

func setupCORS(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

/*	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = "Resource Not Found"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	return
*/
