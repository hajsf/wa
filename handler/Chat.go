package handler

import (
	"fmt"
	"wa/api"
	"wa/utils"

	"go.mau.fi/whatsmeow/types/events"
)

func Chat(evt *events.Message) {
	sender := evt.Info.Chat.User
	pushName := evt.Info.PushName

	msgReceived := evt.Message.GetChat()
	data, _ := utils.PrepareModel(evt.Info.Chat.User,
		sender, pushName, evt.Info.Timestamp.Local().Format("Mon 02-Jan-2006 15:04"),
		evt.Info.ID, "text", fmt.Sprintf("%v", msgReceived), "", "")
	api.Passer.Data <- api.SSEData{
		Event:   "message", // default: source.onmessage = function (event) {}
		Message: data,
	}
}
