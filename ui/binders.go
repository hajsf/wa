package ui

import (
	"context"
	"encoding/json"
	"log"
	"wa/api"
	"wa/global"
	"wa/utils"

	waProto "go.mau.fi/whatsmeow/binary/proto"
	"google.golang.org/protobuf/proto"
)

// go install github.com/jchv/go-webview2@latest

type IncrementResult struct {
	Count uint `json:"count"`
}

func JSbinding() {
	// set title, called from JS as: window.setTitle("home");
	Wv.Bind("setTitle", func(title string) {
		Wv.SetTitle(title)
	})

	// set title, called from JS as: window.send(sender, message);
	Wv.Bind("send", func(sender, reply string) []byte {

		resp := make(map[string]string)
		msg := &waProto.Message{
			Conversation: proto.String(reply),
		}

		jid, ok := utils.ParseJID(sender)
		if !ok {
			resp["error"] = "Wrong number"
			jsonResp, err := json.Marshal(resp)
			if err != nil {
				log.Fatalf("Error happened in JSON marshal. Err: %s", err)
			}
			return jsonResp
		}
		send, err := api.Client.SendMessage(context.Background(), jid, "", msg) // jid = recipient

		if err != nil {
			global.Log.Errorf("Error sending message: %v", err)
		} else {
			global.Log.Infof("Message sent (server timestamp: %s)", send)
		}

		resp["message"] = "Message sent"
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		return jsonResp

	})

	// called at the JavaScript as: window.increment(count).then(result => {...});
	Wv.Bind("increment", func(count uint) IncrementResult {
		count++
		return IncrementResult{Count: count}
	})

}
