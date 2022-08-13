package server

import (
	"encoding/json"
	"net/http"
	"wa/api"
	"wa/global"
	"wa/utils"

	waProto "go.mau.fi/whatsmeow/binary/proto"
	waLog "go.mau.fi/whatsmeow/util/log"
	"google.golang.org/protobuf/proto"
)

type Reply struct {
	Sender  string `json:"sender"`
	Message string `json:"message"`
}

var Log waLog.Logger

func reply(w http.ResponseWriter, r *http.Request) {

	/*	w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	*/
	// Declare a new struct to hold the json data.
	var rply Reply

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&rply)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// fmt.Println("send to:", rply.Sender, ", a reply of:", rply.Message)

	msg := &waProto.Message{
		Conversation: proto.String(rply.Message),
	}

	jid, ok := utils.ParseJID(rply.Sender)
	if !ok {
		return
	}
	send, err := api.Client.SendMessage(global.Ctx, jid, "", msg) // jid = recipient

	if err != nil {
		global.Log.Errorf("Error sending message: %v", err)
	} else {
		global.Log.Infof("Message sent (server timestamp: %s)", send)
	}

	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = "Message sent"
	/*	jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp) */
	json.NewEncoder(w).Encode(resp)
}
